// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/bosh-backup-and-restore/ssh"
)

type FakeSSHConnectionFactory struct {
	Stub        func(host, user, privateKey string) (ssh.SSHConnection, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		host       string
		user       string
		privateKey string
	}
	returns struct {
		result1 ssh.SSHConnection
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 ssh.SSHConnection
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSSHConnectionFactory) Spy(host string, user string, privateKey string) (ssh.SSHConnection, error) {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		host       string
		user       string
		privateKey string
	}{host, user, privateKey})
	fake.recordInvocation("SSHConnectionFactory", []interface{}{host, user, privateKey})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(host, user, privateKey)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.returns.result1, fake.returns.result2
}

func (fake *FakeSSHConnectionFactory) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeSSHConnectionFactory) ArgsForCall(i int) (string, string, string) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].host, fake.argsForCall[i].user, fake.argsForCall[i].privateKey
}

func (fake *FakeSSHConnectionFactory) Returns(result1 ssh.SSHConnection, result2 error) {
	fake.Stub = nil
	fake.returns = struct {
		result1 ssh.SSHConnection
		result2 error
	}{result1, result2}
}

func (fake *FakeSSHConnectionFactory) ReturnsOnCall(i int, result1 ssh.SSHConnection, result2 error) {
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 ssh.SSHConnection
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 ssh.SSHConnection
		result2 error
	}{result1, result2}
}

func (fake *FakeSSHConnectionFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSSHConnectionFactory) recordInvocation(key string, args []interface{}) {
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

var _ ssh.SSHConnectionFactory = new(FakeSSHConnectionFactory).Spy