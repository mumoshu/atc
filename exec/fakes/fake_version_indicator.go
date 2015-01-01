// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/exec"
)

type FakeVersionIndicator struct {
	VersionStub        func() atc.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns struct {
		result1 atc.Version
	}
	MetadataStub        func() []atc.MetadataField
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct{}
	metadataReturns struct {
		result1 []atc.MetadataField
	}
}

func (fake *FakeVersionIndicator) Version() atc.Version {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	} else {
		return fake.versionReturns.result1
	}
}

func (fake *FakeVersionIndicator) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeVersionIndicator) VersionReturns(result1 atc.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeVersionIndicator) Metadata() []atc.MetadataField {
	fake.metadataMutex.Lock()
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct{}{})
	fake.metadataMutex.Unlock()
	if fake.MetadataStub != nil {
		return fake.MetadataStub()
	} else {
		return fake.metadataReturns.result1
	}
}

func (fake *FakeVersionIndicator) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *FakeVersionIndicator) MetadataReturns(result1 []atc.MetadataField) {
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 []atc.MetadataField
	}{result1}
}

var _ exec.VersionIndicator = new(FakeVersionIndicator)
