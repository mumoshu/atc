// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakeNotifier struct {
	NotifyStub        func() <-chan struct{}
	notifyMutex       sync.RWMutex
	notifyArgsForCall []struct{}
	notifyReturns struct {
		result1 <-chan struct{}
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns struct {
		result1 error
	}
}

func (fake *FakeNotifier) Notify() <-chan struct{} {
	fake.notifyMutex.Lock()
	fake.notifyArgsForCall = append(fake.notifyArgsForCall, struct{}{})
	fake.notifyMutex.Unlock()
	if fake.NotifyStub != nil {
		return fake.NotifyStub()
	} else {
		return fake.notifyReturns.result1
	}
}

func (fake *FakeNotifier) NotifyCallCount() int {
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	return len(fake.notifyArgsForCall)
}

func (fake *FakeNotifier) NotifyReturns(result1 <-chan struct{}) {
	fake.NotifyStub = nil
	fake.notifyReturns = struct {
		result1 <-chan struct{}
	}{result1}
}

func (fake *FakeNotifier) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeNotifier) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeNotifier) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

var _ db.Notifier = new(FakeNotifier)
