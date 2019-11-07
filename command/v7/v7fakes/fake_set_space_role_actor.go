// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeSetSpaceRoleActor struct {
	CreateOrgRoleByUserGUIDStub        func(constant.RoleType, string, string) (v7action.Role, v7action.Warnings, error)
	createOrgRoleByUserGUIDMutex       sync.RWMutex
	createOrgRoleByUserGUIDArgsForCall []struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
	}
	createOrgRoleByUserGUIDReturns struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	createOrgRoleByUserGUIDReturnsOnCall map[int]struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	CreateOrgRoleByUserNameStub        func(constant.RoleType, string, string, string) (v7action.Role, v7action.Warnings, error)
	createOrgRoleByUserNameMutex       sync.RWMutex
	createOrgRoleByUserNameArgsForCall []struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
		arg4 string
	}
	createOrgRoleByUserNameReturns struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	createOrgRoleByUserNameReturnsOnCall map[int]struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	CreateSpaceRoleByUserGUIDStub        func(constant.RoleType, string, string) (v7action.Role, v7action.Warnings, error)
	createSpaceRoleByUserGUIDMutex       sync.RWMutex
	createSpaceRoleByUserGUIDArgsForCall []struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
	}
	createSpaceRoleByUserGUIDReturns struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	createSpaceRoleByUserGUIDReturnsOnCall map[int]struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	CreateSpaceRoleByUserNameStub        func(constant.RoleType, string, string, string) (v7action.Role, v7action.Warnings, error)
	createSpaceRoleByUserNameMutex       sync.RWMutex
	createSpaceRoleByUserNameArgsForCall []struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
		arg4 string
	}
	createSpaceRoleByUserNameReturns struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	createSpaceRoleByUserNameReturnsOnCall map[int]struct {
		result1 v7action.Role
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
	GetUserStub        func(string, string) (v7action.User, error)
	getUserMutex       sync.RWMutex
	getUserArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getUserReturns struct {
		result1 v7action.User
		result2 error
	}
	getUserReturnsOnCall map[int]struct {
		result1 v7action.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserGUID(arg1 constant.RoleType, arg2 string, arg3 string) (v7action.Role, v7action.Warnings, error) {
	fake.createOrgRoleByUserGUIDMutex.Lock()
	ret, specificReturn := fake.createOrgRoleByUserGUIDReturnsOnCall[len(fake.createOrgRoleByUserGUIDArgsForCall)]
	fake.createOrgRoleByUserGUIDArgsForCall = append(fake.createOrgRoleByUserGUIDArgsForCall, struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateOrgRoleByUserGUID", []interface{}{arg1, arg2, arg3})
	fake.createOrgRoleByUserGUIDMutex.Unlock()
	if fake.CreateOrgRoleByUserGUIDStub != nil {
		return fake.CreateOrgRoleByUserGUIDStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createOrgRoleByUserGUIDReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserGUIDCallCount() int {
	fake.createOrgRoleByUserGUIDMutex.RLock()
	defer fake.createOrgRoleByUserGUIDMutex.RUnlock()
	return len(fake.createOrgRoleByUserGUIDArgsForCall)
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserGUIDCalls(stub func(constant.RoleType, string, string) (v7action.Role, v7action.Warnings, error)) {
	fake.createOrgRoleByUserGUIDMutex.Lock()
	defer fake.createOrgRoleByUserGUIDMutex.Unlock()
	fake.CreateOrgRoleByUserGUIDStub = stub
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserGUIDArgsForCall(i int) (constant.RoleType, string, string) {
	fake.createOrgRoleByUserGUIDMutex.RLock()
	defer fake.createOrgRoleByUserGUIDMutex.RUnlock()
	argsForCall := fake.createOrgRoleByUserGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserGUIDReturns(result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createOrgRoleByUserGUIDMutex.Lock()
	defer fake.createOrgRoleByUserGUIDMutex.Unlock()
	fake.CreateOrgRoleByUserGUIDStub = nil
	fake.createOrgRoleByUserGUIDReturns = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserGUIDReturnsOnCall(i int, result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createOrgRoleByUserGUIDMutex.Lock()
	defer fake.createOrgRoleByUserGUIDMutex.Unlock()
	fake.CreateOrgRoleByUserGUIDStub = nil
	if fake.createOrgRoleByUserGUIDReturnsOnCall == nil {
		fake.createOrgRoleByUserGUIDReturnsOnCall = make(map[int]struct {
			result1 v7action.Role
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.createOrgRoleByUserGUIDReturnsOnCall[i] = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserName(arg1 constant.RoleType, arg2 string, arg3 string, arg4 string) (v7action.Role, v7action.Warnings, error) {
	fake.createOrgRoleByUserNameMutex.Lock()
	ret, specificReturn := fake.createOrgRoleByUserNameReturnsOnCall[len(fake.createOrgRoleByUserNameArgsForCall)]
	fake.createOrgRoleByUserNameArgsForCall = append(fake.createOrgRoleByUserNameArgsForCall, struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CreateOrgRoleByUserName", []interface{}{arg1, arg2, arg3, arg4})
	fake.createOrgRoleByUserNameMutex.Unlock()
	if fake.CreateOrgRoleByUserNameStub != nil {
		return fake.CreateOrgRoleByUserNameStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createOrgRoleByUserNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserNameCallCount() int {
	fake.createOrgRoleByUserNameMutex.RLock()
	defer fake.createOrgRoleByUserNameMutex.RUnlock()
	return len(fake.createOrgRoleByUserNameArgsForCall)
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserNameCalls(stub func(constant.RoleType, string, string, string) (v7action.Role, v7action.Warnings, error)) {
	fake.createOrgRoleByUserNameMutex.Lock()
	defer fake.createOrgRoleByUserNameMutex.Unlock()
	fake.CreateOrgRoleByUserNameStub = stub
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserNameArgsForCall(i int) (constant.RoleType, string, string, string) {
	fake.createOrgRoleByUserNameMutex.RLock()
	defer fake.createOrgRoleByUserNameMutex.RUnlock()
	argsForCall := fake.createOrgRoleByUserNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserNameReturns(result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createOrgRoleByUserNameMutex.Lock()
	defer fake.createOrgRoleByUserNameMutex.Unlock()
	fake.CreateOrgRoleByUserNameStub = nil
	fake.createOrgRoleByUserNameReturns = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetSpaceRoleActor) CreateOrgRoleByUserNameReturnsOnCall(i int, result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createOrgRoleByUserNameMutex.Lock()
	defer fake.createOrgRoleByUserNameMutex.Unlock()
	fake.CreateOrgRoleByUserNameStub = nil
	if fake.createOrgRoleByUserNameReturnsOnCall == nil {
		fake.createOrgRoleByUserNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Role
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.createOrgRoleByUserNameReturnsOnCall[i] = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserGUID(arg1 constant.RoleType, arg2 string, arg3 string) (v7action.Role, v7action.Warnings, error) {
	fake.createSpaceRoleByUserGUIDMutex.Lock()
	ret, specificReturn := fake.createSpaceRoleByUserGUIDReturnsOnCall[len(fake.createSpaceRoleByUserGUIDArgsForCall)]
	fake.createSpaceRoleByUserGUIDArgsForCall = append(fake.createSpaceRoleByUserGUIDArgsForCall, struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateSpaceRoleByUserGUID", []interface{}{arg1, arg2, arg3})
	fake.createSpaceRoleByUserGUIDMutex.Unlock()
	if fake.CreateSpaceRoleByUserGUIDStub != nil {
		return fake.CreateSpaceRoleByUserGUIDStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createSpaceRoleByUserGUIDReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserGUIDCallCount() int {
	fake.createSpaceRoleByUserGUIDMutex.RLock()
	defer fake.createSpaceRoleByUserGUIDMutex.RUnlock()
	return len(fake.createSpaceRoleByUserGUIDArgsForCall)
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserGUIDCalls(stub func(constant.RoleType, string, string) (v7action.Role, v7action.Warnings, error)) {
	fake.createSpaceRoleByUserGUIDMutex.Lock()
	defer fake.createSpaceRoleByUserGUIDMutex.Unlock()
	fake.CreateSpaceRoleByUserGUIDStub = stub
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserGUIDArgsForCall(i int) (constant.RoleType, string, string) {
	fake.createSpaceRoleByUserGUIDMutex.RLock()
	defer fake.createSpaceRoleByUserGUIDMutex.RUnlock()
	argsForCall := fake.createSpaceRoleByUserGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserGUIDReturns(result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createSpaceRoleByUserGUIDMutex.Lock()
	defer fake.createSpaceRoleByUserGUIDMutex.Unlock()
	fake.CreateSpaceRoleByUserGUIDStub = nil
	fake.createSpaceRoleByUserGUIDReturns = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserGUIDReturnsOnCall(i int, result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createSpaceRoleByUserGUIDMutex.Lock()
	defer fake.createSpaceRoleByUserGUIDMutex.Unlock()
	fake.CreateSpaceRoleByUserGUIDStub = nil
	if fake.createSpaceRoleByUserGUIDReturnsOnCall == nil {
		fake.createSpaceRoleByUserGUIDReturnsOnCall = make(map[int]struct {
			result1 v7action.Role
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.createSpaceRoleByUserGUIDReturnsOnCall[i] = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserName(arg1 constant.RoleType, arg2 string, arg3 string, arg4 string) (v7action.Role, v7action.Warnings, error) {
	fake.createSpaceRoleByUserNameMutex.Lock()
	ret, specificReturn := fake.createSpaceRoleByUserNameReturnsOnCall[len(fake.createSpaceRoleByUserNameArgsForCall)]
	fake.createSpaceRoleByUserNameArgsForCall = append(fake.createSpaceRoleByUserNameArgsForCall, struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CreateSpaceRoleByUserName", []interface{}{arg1, arg2, arg3, arg4})
	fake.createSpaceRoleByUserNameMutex.Unlock()
	if fake.CreateSpaceRoleByUserNameStub != nil {
		return fake.CreateSpaceRoleByUserNameStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createSpaceRoleByUserNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserNameCallCount() int {
	fake.createSpaceRoleByUserNameMutex.RLock()
	defer fake.createSpaceRoleByUserNameMutex.RUnlock()
	return len(fake.createSpaceRoleByUserNameArgsForCall)
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserNameCalls(stub func(constant.RoleType, string, string, string) (v7action.Role, v7action.Warnings, error)) {
	fake.createSpaceRoleByUserNameMutex.Lock()
	defer fake.createSpaceRoleByUserNameMutex.Unlock()
	fake.CreateSpaceRoleByUserNameStub = stub
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserNameArgsForCall(i int) (constant.RoleType, string, string, string) {
	fake.createSpaceRoleByUserNameMutex.RLock()
	defer fake.createSpaceRoleByUserNameMutex.RUnlock()
	argsForCall := fake.createSpaceRoleByUserNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserNameReturns(result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createSpaceRoleByUserNameMutex.Lock()
	defer fake.createSpaceRoleByUserNameMutex.Unlock()
	fake.CreateSpaceRoleByUserNameStub = nil
	fake.createSpaceRoleByUserNameReturns = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetSpaceRoleActor) CreateSpaceRoleByUserNameReturnsOnCall(i int, result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createSpaceRoleByUserNameMutex.Lock()
	defer fake.createSpaceRoleByUserNameMutex.Unlock()
	fake.CreateSpaceRoleByUserNameStub = nil
	if fake.createSpaceRoleByUserNameReturnsOnCall == nil {
		fake.createSpaceRoleByUserNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Role
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.createSpaceRoleByUserNameReturnsOnCall[i] = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetSpaceRoleActor) GetOrganizationByName(arg1 string) (v7action.Organization, v7action.Warnings, error) {
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

func (fake *FakeSetSpaceRoleActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeSetSpaceRoleActor) GetOrganizationByNameCalls(stub func(string) (v7action.Organization, v7action.Warnings, error)) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = stub
}

func (fake *FakeSetSpaceRoleActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	argsForCall := fake.getOrganizationByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSetSpaceRoleActor) GetOrganizationByNameReturns(result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetSpaceRoleActor) GetOrganizationByNameReturnsOnCall(i int, result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
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

func (fake *FakeSetSpaceRoleActor) GetSpaceByNameAndOrganization(arg1 string, arg2 string) (v7action.Space, v7action.Warnings, error) {
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

func (fake *FakeSetSpaceRoleActor) GetSpaceByNameAndOrganizationCallCount() int {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	return len(fake.getSpaceByNameAndOrganizationArgsForCall)
}

func (fake *FakeSetSpaceRoleActor) GetSpaceByNameAndOrganizationCalls(stub func(string, string) (v7action.Space, v7action.Warnings, error)) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = stub
}

func (fake *FakeSetSpaceRoleActor) GetSpaceByNameAndOrganizationArgsForCall(i int) (string, string) {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	argsForCall := fake.getSpaceByNameAndOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSetSpaceRoleActor) GetSpaceByNameAndOrganizationReturns(result1 v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	fake.getSpaceByNameAndOrganizationReturns = struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetSpaceRoleActor) GetSpaceByNameAndOrganizationReturnsOnCall(i int, result1 v7action.Space, result2 v7action.Warnings, result3 error) {
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

func (fake *FakeSetSpaceRoleActor) GetUser(arg1 string, arg2 string) (v7action.User, error) {
	fake.getUserMutex.Lock()
	ret, specificReturn := fake.getUserReturnsOnCall[len(fake.getUserArgsForCall)]
	fake.getUserArgsForCall = append(fake.getUserArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetUser", []interface{}{arg1, arg2})
	fake.getUserMutex.Unlock()
	if fake.GetUserStub != nil {
		return fake.GetUserStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetSpaceRoleActor) GetUserCallCount() int {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return len(fake.getUserArgsForCall)
}

func (fake *FakeSetSpaceRoleActor) GetUserCalls(stub func(string, string) (v7action.User, error)) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = stub
}

func (fake *FakeSetSpaceRoleActor) GetUserArgsForCall(i int) (string, string) {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	argsForCall := fake.getUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSetSpaceRoleActor) GetUserReturns(result1 v7action.User, result2 error) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	fake.getUserReturns = struct {
		result1 v7action.User
		result2 error
	}{result1, result2}
}

func (fake *FakeSetSpaceRoleActor) GetUserReturnsOnCall(i int, result1 v7action.User, result2 error) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	if fake.getUserReturnsOnCall == nil {
		fake.getUserReturnsOnCall = make(map[int]struct {
			result1 v7action.User
			result2 error
		})
	}
	fake.getUserReturnsOnCall[i] = struct {
		result1 v7action.User
		result2 error
	}{result1, result2}
}

func (fake *FakeSetSpaceRoleActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createOrgRoleByUserGUIDMutex.RLock()
	defer fake.createOrgRoleByUserGUIDMutex.RUnlock()
	fake.createOrgRoleByUserNameMutex.RLock()
	defer fake.createOrgRoleByUserNameMutex.RUnlock()
	fake.createSpaceRoleByUserGUIDMutex.RLock()
	defer fake.createSpaceRoleByUserGUIDMutex.RUnlock()
	fake.createSpaceRoleByUserNameMutex.RLock()
	defer fake.createSpaceRoleByUserNameMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSetSpaceRoleActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.SetSpaceRoleActor = new(FakeSetSpaceRoleActor)