// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/pcf-backup-and-restore/backuper"
)

type FakeDeployment struct {
	IsBackupableStub        func() (bool, error)
	isBackupableMutex       sync.RWMutex
	isBackupableArgsForCall []struct{}
	isBackupableReturns     struct {
		result1 bool
		result2 error
	}
	IsRestorableStub        func() (bool, error)
	isRestorableMutex       sync.RWMutex
	isRestorableArgsForCall []struct{}
	isRestorableReturns     struct {
		result1 bool
		result2 error
	}
	PreBackupLockStub        func() error
	preBackupLockMutex       sync.RWMutex
	preBackupLockArgsForCall []struct{}
	preBackupLockReturns     struct {
		result1 error
	}
	BackupStub        func() error
	backupMutex       sync.RWMutex
	backupArgsForCall []struct{}
	backupReturns     struct {
		result1 error
	}
	PostBackupUnlockStub        func() error
	postBackupUnlockMutex       sync.RWMutex
	postBackupUnlockArgsForCall []struct{}
	postBackupUnlockReturns     struct {
		result1 error
	}
	RestoreStub        func() error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct{}
	restoreReturns     struct {
		result1 error
	}
	CopyRemoteBackupToLocalStub        func(backuper.Artifact) error
	copyRemoteBackupToLocalMutex       sync.RWMutex
	copyRemoteBackupToLocalArgsForCall []struct {
		arg1 backuper.Artifact
	}
	copyRemoteBackupToLocalReturns struct {
		result1 error
	}
	CopyLocalBackupToRemoteStub        func(backuper.Artifact) error
	copyLocalBackupToRemoteMutex       sync.RWMutex
	copyLocalBackupToRemoteArgsForCall []struct {
		arg1 backuper.Artifact
	}
	copyLocalBackupToRemoteReturns struct {
		result1 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct{}
	cleanupReturns     struct {
		result1 error
	}
	InstancesStub        func() []backuper.Instance
	instancesMutex       sync.RWMutex
	instancesArgsForCall []struct{}
	instancesReturns     struct {
		result1 []backuper.Instance
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeployment) IsBackupable() (bool, error) {
	fake.isBackupableMutex.Lock()
	fake.isBackupableArgsForCall = append(fake.isBackupableArgsForCall, struct{}{})
	fake.recordInvocation("IsBackupable", []interface{}{})
	fake.isBackupableMutex.Unlock()
	if fake.IsBackupableStub != nil {
		return fake.IsBackupableStub()
	}
	return fake.isBackupableReturns.result1, fake.isBackupableReturns.result2
}

func (fake *FakeDeployment) IsBackupableCallCount() int {
	fake.isBackupableMutex.RLock()
	defer fake.isBackupableMutex.RUnlock()
	return len(fake.isBackupableArgsForCall)
}

func (fake *FakeDeployment) IsBackupableReturns(result1 bool, result2 error) {
	fake.IsBackupableStub = nil
	fake.isBackupableReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) IsRestorable() (bool, error) {
	fake.isRestorableMutex.Lock()
	fake.isRestorableArgsForCall = append(fake.isRestorableArgsForCall, struct{}{})
	fake.recordInvocation("IsRestorable", []interface{}{})
	fake.isRestorableMutex.Unlock()
	if fake.IsRestorableStub != nil {
		return fake.IsRestorableStub()
	}
	return fake.isRestorableReturns.result1, fake.isRestorableReturns.result2
}

func (fake *FakeDeployment) IsRestorableCallCount() int {
	fake.isRestorableMutex.RLock()
	defer fake.isRestorableMutex.RUnlock()
	return len(fake.isRestorableArgsForCall)
}

func (fake *FakeDeployment) IsRestorableReturns(result1 bool, result2 error) {
	fake.IsRestorableStub = nil
	fake.isRestorableReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) PreBackupLock() error {
	fake.preBackupLockMutex.Lock()
	fake.preBackupLockArgsForCall = append(fake.preBackupLockArgsForCall, struct{}{})
	fake.recordInvocation("PreBackupLock", []interface{}{})
	fake.preBackupLockMutex.Unlock()
	if fake.PreBackupLockStub != nil {
		return fake.PreBackupLockStub()
	}
	return fake.preBackupLockReturns.result1
}

func (fake *FakeDeployment) PreBackupLockCallCount() int {
	fake.preBackupLockMutex.RLock()
	defer fake.preBackupLockMutex.RUnlock()
	return len(fake.preBackupLockArgsForCall)
}

func (fake *FakeDeployment) PreBackupLockReturns(result1 error) {
	fake.PreBackupLockStub = nil
	fake.preBackupLockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Backup() error {
	fake.backupMutex.Lock()
	fake.backupArgsForCall = append(fake.backupArgsForCall, struct{}{})
	fake.recordInvocation("Backup", []interface{}{})
	fake.backupMutex.Unlock()
	if fake.BackupStub != nil {
		return fake.BackupStub()
	}
	return fake.backupReturns.result1
}

func (fake *FakeDeployment) BackupCallCount() int {
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	return len(fake.backupArgsForCall)
}

func (fake *FakeDeployment) BackupReturns(result1 error) {
	fake.BackupStub = nil
	fake.backupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) PostBackupUnlock() error {
	fake.postBackupUnlockMutex.Lock()
	fake.postBackupUnlockArgsForCall = append(fake.postBackupUnlockArgsForCall, struct{}{})
	fake.recordInvocation("PostBackupUnlock", []interface{}{})
	fake.postBackupUnlockMutex.Unlock()
	if fake.PostBackupUnlockStub != nil {
		return fake.PostBackupUnlockStub()
	}
	return fake.postBackupUnlockReturns.result1
}

func (fake *FakeDeployment) PostBackupUnlockCallCount() int {
	fake.postBackupUnlockMutex.RLock()
	defer fake.postBackupUnlockMutex.RUnlock()
	return len(fake.postBackupUnlockArgsForCall)
}

func (fake *FakeDeployment) PostBackupUnlockReturns(result1 error) {
	fake.PostBackupUnlockStub = nil
	fake.postBackupUnlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Restore() error {
	fake.restoreMutex.Lock()
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct{}{})
	fake.recordInvocation("Restore", []interface{}{})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		return fake.RestoreStub()
	}
	return fake.restoreReturns.result1
}

func (fake *FakeDeployment) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeDeployment) RestoreReturns(result1 error) {
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) CopyRemoteBackupToLocal(arg1 backuper.Artifact) error {
	fake.copyRemoteBackupToLocalMutex.Lock()
	fake.copyRemoteBackupToLocalArgsForCall = append(fake.copyRemoteBackupToLocalArgsForCall, struct {
		arg1 backuper.Artifact
	}{arg1})
	fake.recordInvocation("CopyRemoteBackupToLocal", []interface{}{arg1})
	fake.copyRemoteBackupToLocalMutex.Unlock()
	if fake.CopyRemoteBackupToLocalStub != nil {
		return fake.CopyRemoteBackupToLocalStub(arg1)
	}
	return fake.copyRemoteBackupToLocalReturns.result1
}

func (fake *FakeDeployment) CopyRemoteBackupToLocalCallCount() int {
	fake.copyRemoteBackupToLocalMutex.RLock()
	defer fake.copyRemoteBackupToLocalMutex.RUnlock()
	return len(fake.copyRemoteBackupToLocalArgsForCall)
}

func (fake *FakeDeployment) CopyRemoteBackupToLocalArgsForCall(i int) backuper.Artifact {
	fake.copyRemoteBackupToLocalMutex.RLock()
	defer fake.copyRemoteBackupToLocalMutex.RUnlock()
	return fake.copyRemoteBackupToLocalArgsForCall[i].arg1
}

func (fake *FakeDeployment) CopyRemoteBackupToLocalReturns(result1 error) {
	fake.CopyRemoteBackupToLocalStub = nil
	fake.copyRemoteBackupToLocalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) CopyLocalBackupToRemote(arg1 backuper.Artifact) error {
	fake.copyLocalBackupToRemoteMutex.Lock()
	fake.copyLocalBackupToRemoteArgsForCall = append(fake.copyLocalBackupToRemoteArgsForCall, struct {
		arg1 backuper.Artifact
	}{arg1})
	fake.recordInvocation("CopyLocalBackupToRemote", []interface{}{arg1})
	fake.copyLocalBackupToRemoteMutex.Unlock()
	if fake.CopyLocalBackupToRemoteStub != nil {
		return fake.CopyLocalBackupToRemoteStub(arg1)
	}
	return fake.copyLocalBackupToRemoteReturns.result1
}

func (fake *FakeDeployment) CopyLocalBackupToRemoteCallCount() int {
	fake.copyLocalBackupToRemoteMutex.RLock()
	defer fake.copyLocalBackupToRemoteMutex.RUnlock()
	return len(fake.copyLocalBackupToRemoteArgsForCall)
}

func (fake *FakeDeployment) CopyLocalBackupToRemoteArgsForCall(i int) backuper.Artifact {
	fake.copyLocalBackupToRemoteMutex.RLock()
	defer fake.copyLocalBackupToRemoteMutex.RUnlock()
	return fake.copyLocalBackupToRemoteArgsForCall[i].arg1
}

func (fake *FakeDeployment) CopyLocalBackupToRemoteReturns(result1 error) {
	fake.CopyLocalBackupToRemoteStub = nil
	fake.copyLocalBackupToRemoteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Cleanup() error {
	fake.cleanupMutex.Lock()
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct{}{})
	fake.recordInvocation("Cleanup", []interface{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	}
	return fake.cleanupReturns.result1
}

func (fake *FakeDeployment) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeDeployment) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Instances() []backuper.Instance {
	fake.instancesMutex.Lock()
	fake.instancesArgsForCall = append(fake.instancesArgsForCall, struct{}{})
	fake.recordInvocation("Instances", []interface{}{})
	fake.instancesMutex.Unlock()
	if fake.InstancesStub != nil {
		return fake.InstancesStub()
	}
	return fake.instancesReturns.result1
}

func (fake *FakeDeployment) InstancesCallCount() int {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return len(fake.instancesArgsForCall)
}

func (fake *FakeDeployment) InstancesReturns(result1 []backuper.Instance) {
	fake.InstancesStub = nil
	fake.instancesReturns = struct {
		result1 []backuper.Instance
	}{result1}
}

func (fake *FakeDeployment) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isBackupableMutex.RLock()
	defer fake.isBackupableMutex.RUnlock()
	fake.isRestorableMutex.RLock()
	defer fake.isRestorableMutex.RUnlock()
	fake.preBackupLockMutex.RLock()
	defer fake.preBackupLockMutex.RUnlock()
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	fake.postBackupUnlockMutex.RLock()
	defer fake.postBackupUnlockMutex.RUnlock()
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	fake.copyRemoteBackupToLocalMutex.RLock()
	defer fake.copyRemoteBackupToLocalMutex.RUnlock()
	fake.copyLocalBackupToRemoteMutex.RLock()
	defer fake.copyLocalBackupToRemoteMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDeployment) recordInvocation(key string, args []interface{}) {
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

var _ backuper.Deployment = new(FakeDeployment)
