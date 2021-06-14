// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/api"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
)

type IdentityProvider struct {
	DefaultIdentityStub        func() view.Identity
	defaultIdentityMutex       sync.RWMutex
	defaultIdentityArgsForCall []struct{}
	defaultIdentityReturns     struct {
		result1 view.Identity
	}
	defaultIdentityReturnsOnCall map[int]struct {
		result1 view.Identity
	}
	IdentityStub        func(label string) view.Identity
	identityMutex       sync.RWMutex
	identityArgsForCall []struct {
		label string
	}
	identityReturns struct {
		result1 view.Identity
	}
	identityReturnsOnCall map[int]struct {
		result1 view.Identity
	}
	AdminsStub        func() []view.Identity
	adminsMutex       sync.RWMutex
	adminsArgsForCall []struct{}
	adminsReturns     struct {
		result1 []view.Identity
	}
	adminsReturnsOnCall map[int]struct {
		result1 []view.Identity
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IdentityProvider) DefaultIdentity() view.Identity {
	fake.defaultIdentityMutex.Lock()
	ret, specificReturn := fake.defaultIdentityReturnsOnCall[len(fake.defaultIdentityArgsForCall)]
	fake.defaultIdentityArgsForCall = append(fake.defaultIdentityArgsForCall, struct{}{})
	fake.recordInvocation("DefaultIdentity", []interface{}{})
	fake.defaultIdentityMutex.Unlock()
	if fake.DefaultIdentityStub != nil {
		return fake.DefaultIdentityStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.defaultIdentityReturns.result1
}

func (fake *IdentityProvider) DefaultIdentityCallCount() int {
	fake.defaultIdentityMutex.RLock()
	defer fake.defaultIdentityMutex.RUnlock()
	return len(fake.defaultIdentityArgsForCall)
}

func (fake *IdentityProvider) DefaultIdentityReturns(result1 view.Identity) {
	fake.DefaultIdentityStub = nil
	fake.defaultIdentityReturns = struct {
		result1 view.Identity
	}{result1}
}

func (fake *IdentityProvider) DefaultIdentityReturnsOnCall(i int, result1 view.Identity) {
	fake.DefaultIdentityStub = nil
	if fake.defaultIdentityReturnsOnCall == nil {
		fake.defaultIdentityReturnsOnCall = make(map[int]struct {
			result1 view.Identity
		})
	}
	fake.defaultIdentityReturnsOnCall[i] = struct {
		result1 view.Identity
	}{result1}
}

func (fake *IdentityProvider) Identity(label string) view.Identity {
	fake.identityMutex.Lock()
	ret, specificReturn := fake.identityReturnsOnCall[len(fake.identityArgsForCall)]
	fake.identityArgsForCall = append(fake.identityArgsForCall, struct {
		label string
	}{label})
	fake.recordInvocation("Identity", []interface{}{label})
	fake.identityMutex.Unlock()
	if fake.IdentityStub != nil {
		return fake.IdentityStub(label)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.identityReturns.result1
}

func (fake *IdentityProvider) IdentityCallCount() int {
	fake.identityMutex.RLock()
	defer fake.identityMutex.RUnlock()
	return len(fake.identityArgsForCall)
}

func (fake *IdentityProvider) IdentityArgsForCall(i int) string {
	fake.identityMutex.RLock()
	defer fake.identityMutex.RUnlock()
	return fake.identityArgsForCall[i].label
}

func (fake *IdentityProvider) IdentityReturns(result1 view.Identity) {
	fake.IdentityStub = nil
	fake.identityReturns = struct {
		result1 view.Identity
	}{result1}
}

func (fake *IdentityProvider) IdentityReturnsOnCall(i int, result1 view.Identity) {
	fake.IdentityStub = nil
	if fake.identityReturnsOnCall == nil {
		fake.identityReturnsOnCall = make(map[int]struct {
			result1 view.Identity
		})
	}
	fake.identityReturnsOnCall[i] = struct {
		result1 view.Identity
	}{result1}
}

func (fake *IdentityProvider) Admins() []view.Identity {
	fake.adminsMutex.Lock()
	ret, specificReturn := fake.adminsReturnsOnCall[len(fake.adminsArgsForCall)]
	fake.adminsArgsForCall = append(fake.adminsArgsForCall, struct{}{})
	fake.recordInvocation("Admins", []interface{}{})
	fake.adminsMutex.Unlock()
	if fake.AdminsStub != nil {
		return fake.AdminsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.adminsReturns.result1
}

func (fake *IdentityProvider) AdminsCallCount() int {
	fake.adminsMutex.RLock()
	defer fake.adminsMutex.RUnlock()
	return len(fake.adminsArgsForCall)
}

func (fake *IdentityProvider) AdminsReturns(result1 []view.Identity) {
	fake.AdminsStub = nil
	fake.adminsReturns = struct {
		result1 []view.Identity
	}{result1}
}

func (fake *IdentityProvider) AdminsReturnsOnCall(i int, result1 []view.Identity) {
	fake.AdminsStub = nil
	if fake.adminsReturnsOnCall == nil {
		fake.adminsReturnsOnCall = make(map[int]struct {
			result1 []view.Identity
		})
	}
	fake.adminsReturnsOnCall[i] = struct {
		result1 []view.Identity
	}{result1}
}

func (fake *IdentityProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.defaultIdentityMutex.RLock()
	defer fake.defaultIdentityMutex.RUnlock()
	fake.identityMutex.RLock()
	defer fake.identityMutex.RUnlock()
	fake.adminsMutex.RLock()
	defer fake.adminsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IdentityProvider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ api.IdentityProvider = new(IdentityProvider)
