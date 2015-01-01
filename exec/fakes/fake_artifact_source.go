// This file was generated by counterfeiter
package fakes

import (
	"io"
	"os"
	"sync"

	"github.com/concourse/atc/exec"
)

type FakeArtifactSource struct {
	RunStub        func(signals <-chan os.Signal, ready chan<- struct{}) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}
	runReturns struct {
		result1 error
	}
	StreamToStub        func(exec.ArtifactDestination) error
	streamToMutex       sync.RWMutex
	streamToArgsForCall []struct {
		arg1 exec.ArtifactDestination
	}
	streamToReturns struct {
		result1 error
	}
	StreamFileStub        func(path string) (io.ReadCloser, error)
	streamFileMutex       sync.RWMutex
	streamFileArgsForCall []struct {
		path string
	}
	streamFileReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	ReleaseStub        func() error
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct{}
	releaseReturns struct {
		result1 error
	}
}

func (fake *FakeArtifactSource) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}{signals, ready})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(signals, ready)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeArtifactSource) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeArtifactSource) RunArgsForCall(i int) (<-chan os.Signal, chan<- struct{}) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].signals, fake.runArgsForCall[i].ready
}

func (fake *FakeArtifactSource) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifactSource) StreamTo(arg1 exec.ArtifactDestination) error {
	fake.streamToMutex.Lock()
	fake.streamToArgsForCall = append(fake.streamToArgsForCall, struct {
		arg1 exec.ArtifactDestination
	}{arg1})
	fake.streamToMutex.Unlock()
	if fake.StreamToStub != nil {
		return fake.StreamToStub(arg1)
	} else {
		return fake.streamToReturns.result1
	}
}

func (fake *FakeArtifactSource) StreamToCallCount() int {
	fake.streamToMutex.RLock()
	defer fake.streamToMutex.RUnlock()
	return len(fake.streamToArgsForCall)
}

func (fake *FakeArtifactSource) StreamToArgsForCall(i int) exec.ArtifactDestination {
	fake.streamToMutex.RLock()
	defer fake.streamToMutex.RUnlock()
	return fake.streamToArgsForCall[i].arg1
}

func (fake *FakeArtifactSource) StreamToReturns(result1 error) {
	fake.StreamToStub = nil
	fake.streamToReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifactSource) StreamFile(path string) (io.ReadCloser, error) {
	fake.streamFileMutex.Lock()
	fake.streamFileArgsForCall = append(fake.streamFileArgsForCall, struct {
		path string
	}{path})
	fake.streamFileMutex.Unlock()
	if fake.StreamFileStub != nil {
		return fake.StreamFileStub(path)
	} else {
		return fake.streamFileReturns.result1, fake.streamFileReturns.result2
	}
}

func (fake *FakeArtifactSource) StreamFileCallCount() int {
	fake.streamFileMutex.RLock()
	defer fake.streamFileMutex.RUnlock()
	return len(fake.streamFileArgsForCall)
}

func (fake *FakeArtifactSource) StreamFileArgsForCall(i int) string {
	fake.streamFileMutex.RLock()
	defer fake.streamFileMutex.RUnlock()
	return fake.streamFileArgsForCall[i].path
}

func (fake *FakeArtifactSource) StreamFileReturns(result1 io.ReadCloser, result2 error) {
	fake.StreamFileStub = nil
	fake.streamFileReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifactSource) Release() error {
	fake.releaseMutex.Lock()
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct{}{})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		return fake.ReleaseStub()
	} else {
		return fake.releaseReturns.result1
	}
}

func (fake *FakeArtifactSource) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeArtifactSource) ReleaseReturns(result1 error) {
	fake.ReleaseStub = nil
	fake.releaseReturns = struct {
		result1 error
	}{result1}
}

var _ exec.ArtifactSource = new(FakeArtifactSource)
