// This file was generated by counterfeiter
package counterfeitertestsfakes

import (
	"sync"

	"github.com/tjarratt/counterfeitertests"
	. "gopkg.in/src-d/go-git.v4"
)

type FakeJustAnExample struct {
	DoesStuffWithGitStub        func(options CloneOptions)
	doesStuffWithGitMutex       sync.RWMutex
	doesStuffWithGitArgsForCall []struct {
		options CloneOptions
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeJustAnExample) DoesStuffWithGit(options CloneOptions) {
	fake.doesStuffWithGitMutex.Lock()
	fake.doesStuffWithGitArgsForCall = append(fake.doesStuffWithGitArgsForCall, struct {
		options CloneOptions
	}{options})
	fake.recordInvocation("DoesStuffWithGit", []interface{}{options})
	fake.doesStuffWithGitMutex.Unlock()
	if fake.DoesStuffWithGitStub != nil {
		fake.DoesStuffWithGitStub(options)
	}
}

func (fake *FakeJustAnExample) DoesStuffWithGitCallCount() int {
	fake.doesStuffWithGitMutex.RLock()
	defer fake.doesStuffWithGitMutex.RUnlock()
	return len(fake.doesStuffWithGitArgsForCall)
}

func (fake *FakeJustAnExample) DoesStuffWithGitArgsForCall(i int) CloneOptions {
	fake.doesStuffWithGitMutex.RLock()
	defer fake.doesStuffWithGitMutex.RUnlock()
	return fake.doesStuffWithGitArgsForCall[i].options
}

func (fake *FakeJustAnExample) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doesStuffWithGitMutex.RLock()
	defer fake.doesStuffWithGitMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeJustAnExample) recordInvocation(key string, args []interface{}) {
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

var _ counterfeitertests.JustAnExample = new(FakeJustAnExample)