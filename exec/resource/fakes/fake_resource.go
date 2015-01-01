// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/exec/resource"
)

type FakeResource struct {
	GetStub        func(atc.Source, atc.Params, atc.Version) resource.VersionedSource
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 atc.Source
		arg2 atc.Params
		arg3 atc.Version
	}
	getReturns struct {
		result1 resource.VersionedSource
	}
	PutStub        func(atc.Source, atc.Params, resource.ArtifactSource) resource.VersionedSource
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 atc.Source
		arg2 atc.Params
		arg3 resource.ArtifactSource
	}
	putReturns struct {
		result1 resource.VersionedSource
	}
	CheckStub        func(atc.Source, atc.Version) ([]atc.Version, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 atc.Source
		arg2 atc.Version
	}
	checkReturns struct {
		result1 []atc.Version
		result2 error
	}
	ReleaseStub        func() error
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct{}
	releaseReturns struct {
		result1 error
	}
}

func (fake *FakeResource) Get(arg1 atc.Source, arg2 atc.Params, arg3 atc.Version) resource.VersionedSource {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 atc.Source
		arg2 atc.Params
		arg3 atc.Version
	}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3)
	} else {
		return fake.getReturns.result1
	}
}

func (fake *FakeResource) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeResource) GetArgsForCall(i int) (atc.Source, atc.Params, atc.Version) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2, fake.getArgsForCall[i].arg3
}

func (fake *FakeResource) GetReturns(result1 resource.VersionedSource) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 resource.VersionedSource
	}{result1}
}

func (fake *FakeResource) Put(arg1 atc.Source, arg2 atc.Params, arg3 resource.ArtifactSource) resource.VersionedSource {
	fake.putMutex.Lock()
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 atc.Source
		arg2 atc.Params
		arg3 resource.ArtifactSource
	}{arg1, arg2, arg3})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2, arg3)
	} else {
		return fake.putReturns.result1
	}
}

func (fake *FakeResource) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeResource) PutArgsForCall(i int) (atc.Source, atc.Params, resource.ArtifactSource) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].arg1, fake.putArgsForCall[i].arg2, fake.putArgsForCall[i].arg3
}

func (fake *FakeResource) PutReturns(result1 resource.VersionedSource) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 resource.VersionedSource
	}{result1}
}

func (fake *FakeResource) Check(arg1 atc.Source, arg2 atc.Version) ([]atc.Version, error) {
	fake.checkMutex.Lock()
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 atc.Source
		arg2 atc.Version
	}{arg1, arg2})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub(arg1, arg2)
	} else {
		return fake.checkReturns.result1, fake.checkReturns.result2
	}
}

func (fake *FakeResource) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeResource) CheckArgsForCall(i int) (atc.Source, atc.Version) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return fake.checkArgsForCall[i].arg1, fake.checkArgsForCall[i].arg2
}

func (fake *FakeResource) CheckReturns(result1 []atc.Version, result2 error) {
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 []atc.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Release() error {
	fake.releaseMutex.Lock()
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct{}{})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		return fake.ReleaseStub()
	} else {
		return fake.releaseReturns.result1
	}
}

func (fake *FakeResource) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeResource) ReleaseReturns(result1 error) {
	fake.ReleaseStub = nil
	fake.releaseReturns = struct {
		result1 error
	}{result1}
}

var _ resource.Resource = new(FakeResource)
