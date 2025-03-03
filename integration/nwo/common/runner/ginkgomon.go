/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
/*
Ginkgomon provides ginkgo test helpers.
*/
package runner

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/flogging"
)

var logger = flogging.MustGetLogger("nwo.runner")

// Config defines a ginkgomon Runner.
type Config struct {
	Command           *exec.Cmd     // process to be executed
	Name              string        // prefixes all output lines
	AnsiColorCode     string        // colors the output
	StartCheck        string        // text to match to indicate successful start.
	StartCheckTimeout time.Duration // how long to wait to see StartCheck
	Cleanup           func()        // invoked once the process exits
	Stdout, Stderr    io.Writer
}

/*
Runner invokes a new process using gomega's gexec package.

If a start check is defined, the runner will wait until it sees the start check
before declaring ready.

Runner implements gexec.Exiter and gbytes.BufferProvider, so you can test exit
codes and process output using the appropriate gomega matchers:
http://onsi.github.io/gomega/#gexec-testing-external-processes
*/
type Runner struct {
	config            Config
	Command           *exec.Cmd
	Name              string
	AnsiColorCode     string
	StartCheck        string
	StartCheckTimeout time.Duration
	Cleanup           func()
	session           *gexec.Session
	sessionReady      chan struct{}
	stop              chan os.Signal
}

// New creates a ginkgomon Runner from a config object. Runners must be created
// with New to properly initialize their internal state.
func New(config Config) *Runner {
	return &Runner{
		config:            config,
		Name:              config.Name,
		Command:           config.Command,
		AnsiColorCode:     config.AnsiColorCode,
		StartCheck:        config.StartCheck,
		StartCheckTimeout: config.StartCheckTimeout,
		Cleanup:           config.Cleanup,
		sessionReady:      make(chan struct{}),
		stop:              make(chan os.Signal),
	}
}

// ExitCode returns the exit code of the process, or -1 if the process has not
// exited.  It can be used with the gexec.Exit matcher.
func (r *Runner) ExitCode() int {
	if r.sessionReady == nil {
		ginkgo.Fail(fmt.Sprintf("ginkgomon.Runner '%s' improperly created without using New", r.Name))
	}
	<-r.sessionReady
	return r.session.ExitCode()
}

// Buffer returns a gbytes.Buffer, for use with the gbytes.Say matcher.
func (r *Runner) Buffer() *gbytes.Buffer {
	if r.sessionReady == nil {
		ginkgo.Fail(fmt.Sprintf("ginkgomon.Runner '%s' improperly created without using New", r.Name))
	}
	<-r.sessionReady
	return r.session.Buffer()
}

// Err returns the gbytes.Buffer associated with the stderr stream.
// For use with the gbytes.Say matcher.
func (r *Runner) Err() *gbytes.Buffer {
	if r.sessionReady == nil {
		ginkgo.Fail(fmt.Sprintf("ginkgomon.Runner '%s' improperly created without using New", r.Name))
	}
	<-r.sessionReady
	return r.session.Err
}

func (r *Runner) Run(sigChan <-chan os.Signal, ready chan<- struct{}) error {
	defer ginkgo.GinkgoRecover()

	var detectStartCheck chan bool

	allOutput := gbytes.NewBuffer()

	debugWriter := gexec.NewPrefixedWriter(
		fmt.Sprintf("\x1b[32m[d]\x1b[%s[%s]\x1b[0m ", r.AnsiColorCode, r.Name),
		ginkgo.GinkgoWriter,
	)

	var session *gexec.Session
	var err error
	if r.config.Stdout != nil || r.config.Stderr != nil {
		logger.Infof("running %s with provided stdout/stderr", r.Name)
		session, err = gexec.Start(
			r.Command,
			io.MultiWriter(allOutput, ginkgo.GinkgoWriter, r.config.Stdout),
			io.MultiWriter(allOutput, ginkgo.GinkgoWriter, r.config.Stderr),
		)
	} else {
		logger.Infof("running %s with ginkgo stdout/stderr", r.Name)
		session, err = gexec.Start(
			r.Command,
			gexec.NewPrefixedWriter(
				fmt.Sprintf("\x1b[32m[o]\x1b[%s[%s]\x1b[0m ", r.AnsiColorCode, r.Name),
				io.MultiWriter(allOutput, ginkgo.GinkgoWriter),
			),
			gexec.NewPrefixedWriter(
				fmt.Sprintf("\x1b[91m[e]\x1b[%s[%s]\x1b[0m ", r.AnsiColorCode, r.Name),
				io.MultiWriter(allOutput, ginkgo.GinkgoWriter),
			),
		)
	}

	Ω(err).ShouldNot(HaveOccurred(), fmt.Sprintf("%s failed to start with err: %s", r.Name, err))

	fmt.Fprintf(debugWriter, "spawned %s (pid: %d) wirh args [%v]\n", r.Command.Path, r.Command.Process.Pid, r.Command.Args)

	r.session = session
	if r.sessionReady != nil {
		close(r.sessionReady)
	}

	startCheckDuration := r.StartCheckTimeout
	if startCheckDuration == 0 {
		startCheckDuration = 5 * time.Second
	}

	var startCheckTimeout <-chan time.Time
	if r.StartCheck != "" {
		startCheckTimeout = time.After(startCheckDuration)
	}

	detectStartCheck = allOutput.Detect(r.StartCheck)

	for {
		select {
		case <-detectStartCheck: // works even with empty string
			allOutput.CancelDetects()
			startCheckTimeout = nil
			detectStartCheck = nil
			close(ready)

		case <-startCheckTimeout:
			// clean up hanging process
			session.Kill().Wait()

			// fail to start
			return fmt.Errorf(
				"did not see %s in command's output within %s. full output:\n\n%s",
				r.StartCheck,
				startCheckDuration,
				string(allOutput.Contents()),
			)

		case signal := <-sigChan:
			session.Signal(signal)

		case <-session.Exited:
			if r.Cleanup != nil {
				r.Cleanup()
			}

			if session.ExitCode() == 0 {
				return nil
			}

			return fmt.Errorf("exit status %d", session.ExitCode())
		case signal := <-r.stop:
			if signal != nil {
				session.Signal(signal)
			}

		}
	}
}

func (r *Runner) Stop() {
	r.stop <- syscall.SIGTERM
}

func (r *Runner) PID() (string, int) {
	return r.Command.Path, r.Command.Process.Pid
}

func (r *Runner) Clone() *Runner {
	c := exec.Command(r.config.Command.Path)
	c.Args = r.config.Command.Args
	c.Env = r.config.Command.Env
	c.Dir = r.config.Command.Dir
	return &Runner{
		config:            r.config,
		Name:              r.config.Name,
		Command:           c,
		AnsiColorCode:     r.config.AnsiColorCode,
		StartCheck:        r.config.StartCheck,
		StartCheckTimeout: r.config.StartCheckTimeout,
		Cleanup:           r.config.Cleanup,
		sessionReady:      make(chan struct{}),
		stop:              make(chan os.Signal),
	}
}
