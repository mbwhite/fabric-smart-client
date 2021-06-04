/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package nochaincode_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hyperledger-labs/fabric-smart-client/integration"
)

func TestEndToEnd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Asset Transfer Secured Agreement (S3)")
}

func StartPort() int {
	return integration.AssetTransferSecuredAgreementS3.StartPortForNode()
}
