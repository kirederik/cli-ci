// Code generated by counterfeiter. DO NOT EDIT.
package rpcfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/plugin/v7/rpc"
)

type FakePluginActor struct {
	GetApplicationsBySpaceStub        func(string) ([]v7action.Application, v7action.Warnings, error)
	getApplicationsBySpaceMutex       sync.RWMutex
	getApplicationsBySpaceArgsForCall []struct {
		arg1 string
	}
	getApplicationsBySpaceReturns struct {
		result1 []v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	getApplicationsBySpaceReturnsOnCall map[int]struct {
		result1 []v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	GetDetailedAppSummaryStub        func(string, string, bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error)
	getDetailedAppSummaryMutex       sync.RWMutex
	getDetailedAppSummaryArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	getDetailedAppSummaryReturns struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}
	getDetailedAppSummaryReturnsOnCall map[int]struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}
	GetOrganizationByNameStub        func(string) (v7action.Organization, v7action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		arg1 string
	}
	getOrganizationByNameReturns struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	GetOrganizationSpacesStub        func(string) ([]v7action.Space, v7action.Warnings, error)
	getOrganizationSpacesMutex       sync.RWMutex
	getOrganizationSpacesArgsForCall []struct {
		arg1 string
	}
	getOrganizationSpacesReturns struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationSpacesReturnsOnCall map[int]struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	GetSpaceByNameAndOrganizationStub        func(string, string) (v7action.Space, v7action.Warnings, error)
	getSpaceByNameAndOrganizationMutex       sync.RWMutex
	getSpaceByNameAndOrganizationArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceByNameAndOrganizationReturns struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	getSpaceByNameAndOrganizationReturnsOnCall map[int]struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	RefreshAccessTokenStub        func() (string, error)
	refreshAccessTokenMutex       sync.RWMutex
	refreshAccessTokenArgsForCall []struct {
	}
	refreshAccessTokenReturns struct {
		result1 string
		result2 error
	}
	refreshAccessTokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePluginActor) GetApplicationsBySpace(arg1 string) ([]v7action.Application, v7action.Warnings, error) {
	fake.getApplicationsBySpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationsBySpaceReturnsOnCall[len(fake.getApplicationsBySpaceArgsForCall)]
	fake.getApplicationsBySpaceArgsForCall = append(fake.getApplicationsBySpaceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetApplicationsBySpace", []interface{}{arg1})
	fake.getApplicationsBySpaceMutex.Unlock()
	if fake.GetApplicationsBySpaceStub != nil {
		return fake.GetApplicationsBySpaceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationsBySpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePluginActor) GetApplicationsBySpaceCallCount() int {
	fake.getApplicationsBySpaceMutex.RLock()
	defer fake.getApplicationsBySpaceMutex.RUnlock()
	return len(fake.getApplicationsBySpaceArgsForCall)
}

func (fake *FakePluginActor) GetApplicationsBySpaceCalls(stub func(string) ([]v7action.Application, v7action.Warnings, error)) {
	fake.getApplicationsBySpaceMutex.Lock()
	defer fake.getApplicationsBySpaceMutex.Unlock()
	fake.GetApplicationsBySpaceStub = stub
}

func (fake *FakePluginActor) GetApplicationsBySpaceArgsForCall(i int) string {
	fake.getApplicationsBySpaceMutex.RLock()
	defer fake.getApplicationsBySpaceMutex.RUnlock()
	argsForCall := fake.getApplicationsBySpaceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePluginActor) GetApplicationsBySpaceReturns(result1 []v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationsBySpaceMutex.Lock()
	defer fake.getApplicationsBySpaceMutex.Unlock()
	fake.GetApplicationsBySpaceStub = nil
	fake.getApplicationsBySpaceReturns = struct {
		result1 []v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePluginActor) GetApplicationsBySpaceReturnsOnCall(i int, result1 []v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationsBySpaceMutex.Lock()
	defer fake.getApplicationsBySpaceMutex.Unlock()
	fake.GetApplicationsBySpaceStub = nil
	if fake.getApplicationsBySpaceReturnsOnCall == nil {
		fake.getApplicationsBySpaceReturnsOnCall = make(map[int]struct {
			result1 []v7action.Application
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationsBySpaceReturnsOnCall[i] = struct {
		result1 []v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePluginActor) GetDetailedAppSummary(arg1 string, arg2 string, arg3 bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error) {
	fake.getDetailedAppSummaryMutex.Lock()
	ret, specificReturn := fake.getDetailedAppSummaryReturnsOnCall[len(fake.getDetailedAppSummaryArgsForCall)]
	fake.getDetailedAppSummaryArgsForCall = append(fake.getDetailedAppSummaryArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetDetailedAppSummary", []interface{}{arg1, arg2, arg3})
	fake.getDetailedAppSummaryMutex.Unlock()
	if fake.GetDetailedAppSummaryStub != nil {
		return fake.GetDetailedAppSummaryStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getDetailedAppSummaryReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePluginActor) GetDetailedAppSummaryCallCount() int {
	fake.getDetailedAppSummaryMutex.RLock()
	defer fake.getDetailedAppSummaryMutex.RUnlock()
	return len(fake.getDetailedAppSummaryArgsForCall)
}

func (fake *FakePluginActor) GetDetailedAppSummaryCalls(stub func(string, string, bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error)) {
	fake.getDetailedAppSummaryMutex.Lock()
	defer fake.getDetailedAppSummaryMutex.Unlock()
	fake.GetDetailedAppSummaryStub = stub
}

func (fake *FakePluginActor) GetDetailedAppSummaryArgsForCall(i int) (string, string, bool) {
	fake.getDetailedAppSummaryMutex.RLock()
	defer fake.getDetailedAppSummaryMutex.RUnlock()
	argsForCall := fake.getDetailedAppSummaryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePluginActor) GetDetailedAppSummaryReturns(result1 v7action.DetailedApplicationSummary, result2 v7action.Warnings, result3 error) {
	fake.getDetailedAppSummaryMutex.Lock()
	defer fake.getDetailedAppSummaryMutex.Unlock()
	fake.GetDetailedAppSummaryStub = nil
	fake.getDetailedAppSummaryReturns = struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePluginActor) GetDetailedAppSummaryReturnsOnCall(i int, result1 v7action.DetailedApplicationSummary, result2 v7action.Warnings, result3 error) {
	fake.getDetailedAppSummaryMutex.Lock()
	defer fake.getDetailedAppSummaryMutex.Unlock()
	fake.GetDetailedAppSummaryStub = nil
	if fake.getDetailedAppSummaryReturnsOnCall == nil {
		fake.getDetailedAppSummaryReturnsOnCall = make(map[int]struct {
			result1 v7action.DetailedApplicationSummary
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getDetailedAppSummaryReturnsOnCall[i] = struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePluginActor) GetOrganizationByName(arg1 string) (v7action.Organization, v7action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationByName", []interface{}{arg1})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePluginActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakePluginActor) GetOrganizationByNameCalls(stub func(string) (v7action.Organization, v7action.Warnings, error)) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = stub
}

func (fake *FakePluginActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	argsForCall := fake.getOrganizationByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePluginActor) GetOrganizationByNameReturns(result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePluginActor) GetOrganizationByNameReturnsOnCall(i int, result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Organization
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePluginActor) GetOrganizationSpaces(arg1 string) ([]v7action.Space, v7action.Warnings, error) {
	fake.getOrganizationSpacesMutex.Lock()
	ret, specificReturn := fake.getOrganizationSpacesReturnsOnCall[len(fake.getOrganizationSpacesArgsForCall)]
	fake.getOrganizationSpacesArgsForCall = append(fake.getOrganizationSpacesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationSpaces", []interface{}{arg1})
	fake.getOrganizationSpacesMutex.Unlock()
	if fake.GetOrganizationSpacesStub != nil {
		return fake.GetOrganizationSpacesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationSpacesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePluginActor) GetOrganizationSpacesCallCount() int {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	return len(fake.getOrganizationSpacesArgsForCall)
}

func (fake *FakePluginActor) GetOrganizationSpacesCalls(stub func(string) ([]v7action.Space, v7action.Warnings, error)) {
	fake.getOrganizationSpacesMutex.Lock()
	defer fake.getOrganizationSpacesMutex.Unlock()
	fake.GetOrganizationSpacesStub = stub
}

func (fake *FakePluginActor) GetOrganizationSpacesArgsForCall(i int) string {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	argsForCall := fake.getOrganizationSpacesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePluginActor) GetOrganizationSpacesReturns(result1 []v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationSpacesMutex.Lock()
	defer fake.getOrganizationSpacesMutex.Unlock()
	fake.GetOrganizationSpacesStub = nil
	fake.getOrganizationSpacesReturns = struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePluginActor) GetOrganizationSpacesReturnsOnCall(i int, result1 []v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationSpacesMutex.Lock()
	defer fake.getOrganizationSpacesMutex.Unlock()
	fake.GetOrganizationSpacesStub = nil
	if fake.getOrganizationSpacesReturnsOnCall == nil {
		fake.getOrganizationSpacesReturnsOnCall = make(map[int]struct {
			result1 []v7action.Space
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationSpacesReturnsOnCall[i] = struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePluginActor) GetSpaceByNameAndOrganization(arg1 string, arg2 string) (v7action.Space, v7action.Warnings, error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	ret, specificReturn := fake.getSpaceByNameAndOrganizationReturnsOnCall[len(fake.getSpaceByNameAndOrganizationArgsForCall)]
	fake.getSpaceByNameAndOrganizationArgsForCall = append(fake.getSpaceByNameAndOrganizationArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSpaceByNameAndOrganization", []interface{}{arg1, arg2})
	fake.getSpaceByNameAndOrganizationMutex.Unlock()
	if fake.GetSpaceByNameAndOrganizationStub != nil {
		return fake.GetSpaceByNameAndOrganizationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceByNameAndOrganizationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePluginActor) GetSpaceByNameAndOrganizationCallCount() int {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	return len(fake.getSpaceByNameAndOrganizationArgsForCall)
}

func (fake *FakePluginActor) GetSpaceByNameAndOrganizationCalls(stub func(string, string) (v7action.Space, v7action.Warnings, error)) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = stub
}

func (fake *FakePluginActor) GetSpaceByNameAndOrganizationArgsForCall(i int) (string, string) {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	argsForCall := fake.getSpaceByNameAndOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePluginActor) GetSpaceByNameAndOrganizationReturns(result1 v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	fake.getSpaceByNameAndOrganizationReturns = struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePluginActor) GetSpaceByNameAndOrganizationReturnsOnCall(i int, result1 v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	if fake.getSpaceByNameAndOrganizationReturnsOnCall == nil {
		fake.getSpaceByNameAndOrganizationReturnsOnCall = make(map[int]struct {
			result1 v7action.Space
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSpaceByNameAndOrganizationReturnsOnCall[i] = struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePluginActor) RefreshAccessToken() (string, error) {
	fake.refreshAccessTokenMutex.Lock()
	ret, specificReturn := fake.refreshAccessTokenReturnsOnCall[len(fake.refreshAccessTokenArgsForCall)]
	fake.refreshAccessTokenArgsForCall = append(fake.refreshAccessTokenArgsForCall, struct {
	}{})
	fake.recordInvocation("RefreshAccessToken", []interface{}{})
	fake.refreshAccessTokenMutex.Unlock()
	if fake.RefreshAccessTokenStub != nil {
		return fake.RefreshAccessTokenStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.refreshAccessTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePluginActor) RefreshAccessTokenCallCount() int {
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	return len(fake.refreshAccessTokenArgsForCall)
}

func (fake *FakePluginActor) RefreshAccessTokenCalls(stub func() (string, error)) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = stub
}

func (fake *FakePluginActor) RefreshAccessTokenReturns(result1 string, result2 error) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = nil
	fake.refreshAccessTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakePluginActor) RefreshAccessTokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = nil
	if fake.refreshAccessTokenReturnsOnCall == nil {
		fake.refreshAccessTokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.refreshAccessTokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakePluginActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationsBySpaceMutex.RLock()
	defer fake.getApplicationsBySpaceMutex.RUnlock()
	fake.getDetailedAppSummaryMutex.RLock()
	defer fake.getDetailedAppSummaryMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePluginActor) recordInvocation(key string, args []interface{}) {
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

var _ rpc.PluginActor = new(FakePluginActor)
