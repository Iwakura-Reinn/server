// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/bangumi/server/internal/auth"

	mock "github.com/stretchr/testify/mock"

	session "github.com/bangumi/server/web/session"
)

// SessionManager is an autogenerated mock type for the Manager type
type SessionManager struct {
	mock.Mock
}

type SessionManager_Expecter struct {
	mock *mock.Mock
}

func (_m *SessionManager) EXPECT() *SessionManager_Expecter {
	return &SessionManager_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, a
func (_m *SessionManager) Create(ctx context.Context, a auth.Auth) (string, session.Session, error) {
	ret := _m.Called(ctx, a)

	var r0 string
	var r1 session.Session
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, auth.Auth) (string, session.Session, error)); ok {
		return rf(ctx, a)
	}
	if rf, ok := ret.Get(0).(func(context.Context, auth.Auth) string); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, auth.Auth) session.Session); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Get(1).(session.Session)
	}

	if rf, ok := ret.Get(2).(func(context.Context, auth.Auth) error); ok {
		r2 = rf(ctx, a)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SessionManager_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type SessionManager_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - a auth.Auth
func (_e *SessionManager_Expecter) Create(ctx interface{}, a interface{}) *SessionManager_Create_Call {
	return &SessionManager_Create_Call{Call: _e.mock.On("Create", ctx, a)}
}

func (_c *SessionManager_Create_Call) Run(run func(ctx context.Context, a auth.Auth)) *SessionManager_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(auth.Auth))
	})
	return _c
}

func (_c *SessionManager_Create_Call) Return(_a0 string, _a1 session.Session, _a2 error) *SessionManager_Create_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *SessionManager_Create_Call) RunAndReturn(run func(context.Context, auth.Auth) (string, session.Session, error)) *SessionManager_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key
func (_m *SessionManager) Get(ctx context.Context, key string) (session.Session, error) {
	ret := _m.Called(ctx, key)

	var r0 session.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (session.Session, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) session.Session); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(session.Session)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SessionManager_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type SessionManager_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *SessionManager_Expecter) Get(ctx interface{}, key interface{}) *SessionManager_Get_Call {
	return &SessionManager_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *SessionManager_Get_Call) Run(run func(ctx context.Context, key string)) *SessionManager_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *SessionManager_Get_Call) Return(_a0 session.Session, _a1 error) *SessionManager_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SessionManager_Get_Call) RunAndReturn(run func(context.Context, string) (session.Session, error)) *SessionManager_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Revoke provides a mock function with given fields: ctx, key
func (_m *SessionManager) Revoke(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionManager_Revoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Revoke'
type SessionManager_Revoke_Call struct {
	*mock.Call
}

// Revoke is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *SessionManager_Expecter) Revoke(ctx interface{}, key interface{}) *SessionManager_Revoke_Call {
	return &SessionManager_Revoke_Call{Call: _e.mock.On("Revoke", ctx, key)}
}

func (_c *SessionManager_Revoke_Call) Run(run func(ctx context.Context, key string)) *SessionManager_Revoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *SessionManager_Revoke_Call) Return(_a0 error) *SessionManager_Revoke_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionManager_Revoke_Call) RunAndReturn(run func(context.Context, string) error) *SessionManager_Revoke_Call {
	_c.Call.Return(run)
	return _c
}

// RevokeUser provides a mock function with given fields: ctx, id
func (_m *SessionManager) RevokeUser(ctx context.Context, id uint32) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionManager_RevokeUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevokeUser'
type SessionManager_RevokeUser_Call struct {
	*mock.Call
}

// RevokeUser is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
func (_e *SessionManager_Expecter) RevokeUser(ctx interface{}, id interface{}) *SessionManager_RevokeUser_Call {
	return &SessionManager_RevokeUser_Call{Call: _e.mock.On("RevokeUser", ctx, id)}
}

func (_c *SessionManager_RevokeUser_Call) Run(run func(ctx context.Context, id uint32)) *SessionManager_RevokeUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *SessionManager_RevokeUser_Call) Return(_a0 error) *SessionManager_RevokeUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionManager_RevokeUser_Call) RunAndReturn(run func(context.Context, uint32) error) *SessionManager_RevokeUser_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSessionManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewSessionManager creates a new instance of SessionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSessionManager(t mockConstructorTestingTNewSessionManager) *SessionManager {
	mock := &SessionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
