// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	ldap "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/ldap"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockLdapClientService is an autogenerated mock type for the ClientService type
type MockLdapClientService struct {
	mock.Mock
}

// ImportLdapUser provides a mock function with given fields: params, authInfo
func (_m *MockLdapClientService) ImportLdapUser(params *ldap.ImportLdapUserParams, authInfo runtime.ClientAuthInfoWriter) (*ldap.ImportLdapUserOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *ldap.ImportLdapUserOK
	if rf, ok := ret.Get(0).(func(*ldap.ImportLdapUserParams, runtime.ClientAuthInfoWriter) *ldap.ImportLdapUserOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ldap.ImportLdapUserOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ldap.ImportLdapUserParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PingLdap provides a mock function with given fields: params, authInfo
func (_m *MockLdapClientService) PingLdap(params *ldap.PingLdapParams, authInfo runtime.ClientAuthInfoWriter) (*ldap.PingLdapOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *ldap.PingLdapOK
	if rf, ok := ret.Get(0).(func(*ldap.PingLdapParams, runtime.ClientAuthInfoWriter) *ldap.PingLdapOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ldap.PingLdapOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ldap.PingLdapParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchLdapGroup provides a mock function with given fields: params, authInfo
func (_m *MockLdapClientService) SearchLdapGroup(params *ldap.SearchLdapGroupParams, authInfo runtime.ClientAuthInfoWriter) (*ldap.SearchLdapGroupOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *ldap.SearchLdapGroupOK
	if rf, ok := ret.Get(0).(func(*ldap.SearchLdapGroupParams, runtime.ClientAuthInfoWriter) *ldap.SearchLdapGroupOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ldap.SearchLdapGroupOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ldap.SearchLdapGroupParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchLdapUser provides a mock function with given fields: params, authInfo
func (_m *MockLdapClientService) SearchLdapUser(params *ldap.SearchLdapUserParams, authInfo runtime.ClientAuthInfoWriter) (*ldap.SearchLdapUserOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *ldap.SearchLdapUserOK
	if rf, ok := ret.Get(0).(func(*ldap.SearchLdapUserParams, runtime.ClientAuthInfoWriter) *ldap.SearchLdapUserOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ldap.SearchLdapUserOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ldap.SearchLdapUserParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockLdapClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

type mockConstructorTestingTNewMockLdapClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockLdapClientService creates a new instance of MockLdapClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockLdapClientService(t mockConstructorTestingTNewMockLdapClientService) *MockLdapClientService {
	mock := &MockLdapClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
