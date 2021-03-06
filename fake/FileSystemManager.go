// This file was generated by counterfeiter
package fake

import (
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/phogolabs/parcello"
)

type FileSystemManager struct {
	OpenStub        func(name string) (http.File, error)
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		name string
	}
	openReturns struct {
		result1 http.File
		result2 error
	}
	WalkStub        func(dir string, fn filepath.WalkFunc) error
	walkMutex       sync.RWMutex
	walkArgsForCall []struct {
		dir string
		fn  filepath.WalkFunc
	}
	walkReturns struct {
		result1 error
	}
	OpenFileStub        func(name string, flag int, perm os.FileMode) (parcello.File, error)
	openFileMutex       sync.RWMutex
	openFileArgsForCall []struct {
		name string
		flag int
		perm os.FileMode
	}
	openFileReturns struct {
		result1 parcello.File
		result2 error
	}
	DirStub        func(name string) (parcello.FileSystemManager, error)
	dirMutex       sync.RWMutex
	dirArgsForCall []struct {
		name string
	}
	dirReturns struct {
		result1 parcello.FileSystemManager
		result2 error
	}
	AddStub        func(resource *parcello.Resource) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		resource *parcello.Resource
	}
	addReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FileSystemManager) Open(name string) (http.File, error) {
	fake.openMutex.Lock()
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Open", []interface{}{name})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(name)
	}
	return fake.openReturns.result1, fake.openReturns.result2
}

func (fake *FileSystemManager) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FileSystemManager) OpenArgsForCall(i int) string {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return fake.openArgsForCall[i].name
}

func (fake *FileSystemManager) OpenReturns(result1 http.File, result2 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 http.File
		result2 error
	}{result1, result2}
}

func (fake *FileSystemManager) Walk(dir string, fn filepath.WalkFunc) error {
	fake.walkMutex.Lock()
	fake.walkArgsForCall = append(fake.walkArgsForCall, struct {
		dir string
		fn  filepath.WalkFunc
	}{dir, fn})
	fake.recordInvocation("Walk", []interface{}{dir, fn})
	fake.walkMutex.Unlock()
	if fake.WalkStub != nil {
		return fake.WalkStub(dir, fn)
	}
	return fake.walkReturns.result1
}

func (fake *FileSystemManager) WalkCallCount() int {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	return len(fake.walkArgsForCall)
}

func (fake *FileSystemManager) WalkArgsForCall(i int) (string, filepath.WalkFunc) {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	return fake.walkArgsForCall[i].dir, fake.walkArgsForCall[i].fn
}

func (fake *FileSystemManager) WalkReturns(result1 error) {
	fake.WalkStub = nil
	fake.walkReturns = struct {
		result1 error
	}{result1}
}

func (fake *FileSystemManager) OpenFile(name string, flag int, perm os.FileMode) (parcello.File, error) {
	fake.openFileMutex.Lock()
	fake.openFileArgsForCall = append(fake.openFileArgsForCall, struct {
		name string
		flag int
		perm os.FileMode
	}{name, flag, perm})
	fake.recordInvocation("OpenFile", []interface{}{name, flag, perm})
	fake.openFileMutex.Unlock()
	if fake.OpenFileStub != nil {
		return fake.OpenFileStub(name, flag, perm)
	}
	return fake.openFileReturns.result1, fake.openFileReturns.result2
}

func (fake *FileSystemManager) OpenFileCallCount() int {
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return len(fake.openFileArgsForCall)
}

func (fake *FileSystemManager) OpenFileArgsForCall(i int) (string, int, os.FileMode) {
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return fake.openFileArgsForCall[i].name, fake.openFileArgsForCall[i].flag, fake.openFileArgsForCall[i].perm
}

func (fake *FileSystemManager) OpenFileReturns(result1 parcello.File, result2 error) {
	fake.OpenFileStub = nil
	fake.openFileReturns = struct {
		result1 parcello.File
		result2 error
	}{result1, result2}
}

func (fake *FileSystemManager) Dir(name string) (parcello.FileSystemManager, error) {
	fake.dirMutex.Lock()
	fake.dirArgsForCall = append(fake.dirArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Dir", []interface{}{name})
	fake.dirMutex.Unlock()
	if fake.DirStub != nil {
		return fake.DirStub(name)
	}
	return fake.dirReturns.result1, fake.dirReturns.result2
}

func (fake *FileSystemManager) DirCallCount() int {
	fake.dirMutex.RLock()
	defer fake.dirMutex.RUnlock()
	return len(fake.dirArgsForCall)
}

func (fake *FileSystemManager) DirArgsForCall(i int) string {
	fake.dirMutex.RLock()
	defer fake.dirMutex.RUnlock()
	return fake.dirArgsForCall[i].name
}

func (fake *FileSystemManager) DirReturns(result1 parcello.FileSystemManager, result2 error) {
	fake.DirStub = nil
	fake.dirReturns = struct {
		result1 parcello.FileSystemManager
		result2 error
	}{result1, result2}
}

func (fake *FileSystemManager) Add(resource *parcello.Resource) error {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		resource *parcello.Resource
	}{resource})
	fake.recordInvocation("Add", []interface{}{resource})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(resource)
	}
	return fake.addReturns.result1
}

func (fake *FileSystemManager) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FileSystemManager) AddArgsForCall(i int) *parcello.Resource {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].resource
}

func (fake *FileSystemManager) AddReturns(result1 error) {
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *FileSystemManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	fake.dirMutex.RLock()
	defer fake.dirMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.invocations
}

func (fake *FileSystemManager) recordInvocation(key string, args []interface{}) {
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

var _ parcello.FileSystemManager = new(FileSystemManager)
