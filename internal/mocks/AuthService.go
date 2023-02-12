// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/bangumi/server/internal/auth"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// AuthService is an autogenerated mock type for the Service type
type AuthService struct {
	mock.Mock
}

type AuthService_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthService) EXPECT() *AuthService_Expecter {
	return &AuthService_Expecter{mock: &_m.Mock}
}

// ComparePassword provides a mock function with given fields: hashed, password
func (_m *AuthService) ComparePassword(hashed []byte, password string) (bool, error) {
	ret := _m.Called(hashed, password)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]byte, string) bool); ok {
		r0 = rf(hashed, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(hashed, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_ComparePassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ComparePassword'
type AuthService_ComparePassword_Call struct {
	*mock.Call
}

// ComparePassword is a helper method to define mock.On call
//   - hashed []byte
//   - password string
func (_e *AuthService_Expecter) ComparePassword(hashed interface{}, password interface{}) *AuthService_ComparePassword_Call {
	return &AuthService_ComparePassword_Call{Call: _e.mock.On("ComparePassword", hashed, password)}
}

func (_c *AuthService_ComparePassword_Call) Run(run func(hashed []byte, password string)) *AuthService_ComparePassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(string))
	})
	return _c
}

func (_c *AuthService_ComparePassword_Call) Return(_a0 bool, _a1 error) *AuthService_ComparePassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CreateAccessToken provides a mock function with given fields: ctx, userID, name, expiration
func (_m *AuthService) CreateAccessToken(ctx context.Context, userID uint32, name string, expiration time.Duration) (string, error) {
	ret := _m.Called(ctx, userID, name, expiration)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, uint32, string, time.Duration) string); ok {
		r0 = rf(ctx, userID, name, expiration)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, string, time.Duration) error); ok {
		r1 = rf(ctx, userID, name, expiration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_CreateAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccessToken'
type AuthService_CreateAccessToken_Call struct {
	*mock.Call
}

// CreateAccessToken is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - name string
//   - expiration time.Duration
func (_e *AuthService_Expecter) CreateAccessToken(ctx interface{}, userID interface{}, name interface{}, expiration interface{}) *AuthService_CreateAccessToken_Call {
	return &AuthService_CreateAccessToken_Call{Call: _e.mock.On("CreateAccessToken", ctx, userID, name, expiration)}
}

func (_c *AuthService_CreateAccessToken_Call) Run(run func(ctx context.Context, userID uint32, name string, expiration time.Duration)) *AuthService_CreateAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(string), args[3].(time.Duration))
	})
	return _c
}

func (_c *AuthService_CreateAccessToken_Call) Return(token string, err error) *AuthService_CreateAccessToken_Call {
	_c.Call.Return(token, err)
	return _c
}

// DeleteAccessToken provides a mock function with given fields: ctx, tokenID
func (_m *AuthService) DeleteAccessToken(ctx context.Context, tokenID uint32) (bool, error) {
	ret := _m.Called(ctx, tokenID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, uint32) bool); ok {
		r0 = rf(ctx, tokenID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, tokenID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_DeleteAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAccessToken'
type AuthService_DeleteAccessToken_Call struct {
	*mock.Call
}

// DeleteAccessToken is a helper method to define mock.On call
//   - ctx context.Context
//   - tokenID uint32
func (_e *AuthService_Expecter) DeleteAccessToken(ctx interface{}, tokenID interface{}) *AuthService_DeleteAccessToken_Call {
	return &AuthService_DeleteAccessToken_Call{Call: _e.mock.On("DeleteAccessToken", ctx, tokenID)}
}

func (_c *AuthService_DeleteAccessToken_Call) Run(run func(ctx context.Context, tokenID uint32)) *AuthService_DeleteAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *AuthService_DeleteAccessToken_Call) Return(_a0 bool, _a1 error) *AuthService_DeleteAccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByID provides a mock function with given fields: ctx, userID
func (_m *AuthService) GetByID(ctx context.Context, userID uint32) (auth.Auth, error) {
	ret := _m.Called(ctx, userID)

	var r0 auth.Auth
	if rf, ok := ret.Get(0).(func(context.Context, uint32) auth.Auth); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(auth.Auth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type AuthService_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
func (_e *AuthService_Expecter) GetByID(ctx interface{}, userID interface{}) *AuthService_GetByID_Call {
	return &AuthService_GetByID_Call{Call: _e.mock.On("GetByID", ctx, userID)}
}

func (_c *AuthService_GetByID_Call) Run(run func(ctx context.Context, userID uint32)) *AuthService_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *AuthService_GetByID_Call) Return(_a0 auth.Auth, _a1 error) *AuthService_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByToken provides a mock function with given fields: ctx, token
func (_m *AuthService) GetByToken(ctx context.Context, token string) (auth.Auth, error) {
	ret := _m.Called(ctx, token)

	var r0 auth.Auth
	if rf, ok := ret.Get(0).(func(context.Context, string) auth.Auth); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(auth.Auth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_GetByToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByToken'
type AuthService_GetByToken_Call struct {
	*mock.Call
}

// GetByToken is a helper method to define mock.On call
//   - ctx context.Context
//   - token string
func (_e *AuthService_Expecter) GetByToken(ctx interface{}, token interface{}) *AuthService_GetByToken_Call {
	return &AuthService_GetByToken_Call{Call: _e.mock.On("GetByToken", ctx, token)}
}

func (_c *AuthService_GetByToken_Call) Run(run func(ctx context.Context, token string)) *AuthService_GetByToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthService_GetByToken_Call) Return(_a0 auth.Auth, _a1 error) *AuthService_GetByToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTokenByID provides a mock function with given fields: ctx, tokenID
func (_m *AuthService) GetTokenByID(ctx context.Context, tokenID uint32) (auth.AccessToken, error) {
	ret := _m.Called(ctx, tokenID)

	var r0 auth.AccessToken
	if rf, ok := ret.Get(0).(func(context.Context, uint32) auth.AccessToken); ok {
		r0 = rf(ctx, tokenID)
	} else {
		r0 = ret.Get(0).(auth.AccessToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, tokenID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_GetTokenByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokenByID'
type AuthService_GetTokenByID_Call struct {
	*mock.Call
}

// GetTokenByID is a helper method to define mock.On call
//   - ctx context.Context
//   - tokenID uint32
func (_e *AuthService_Expecter) GetTokenByID(ctx interface{}, tokenID interface{}) *AuthService_GetTokenByID_Call {
	return &AuthService_GetTokenByID_Call{Call: _e.mock.On("GetTokenByID", ctx, tokenID)}
}

func (_c *AuthService_GetTokenByID_Call) Run(run func(ctx context.Context, tokenID uint32)) *AuthService_GetTokenByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *AuthService_GetTokenByID_Call) Return(_a0 auth.AccessToken, _a1 error) *AuthService_GetTokenByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListAccessToken provides a mock function with given fields: ctx, userID
func (_m *AuthService) ListAccessToken(ctx context.Context, userID uint32) ([]auth.AccessToken, error) {
	ret := _m.Called(ctx, userID)

	var r0 []auth.AccessToken
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []auth.AccessToken); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]auth.AccessToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_ListAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAccessToken'
type AuthService_ListAccessToken_Call struct {
	*mock.Call
}

// ListAccessToken is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
func (_e *AuthService_Expecter) ListAccessToken(ctx interface{}, userID interface{}) *AuthService_ListAccessToken_Call {
	return &AuthService_ListAccessToken_Call{Call: _e.mock.On("ListAccessToken", ctx, userID)}
}

func (_c *AuthService_ListAccessToken_Call) Run(run func(ctx context.Context, userID uint32)) *AuthService_ListAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *AuthService_ListAccessToken_Call) Return(_a0 []auth.AccessToken, _a1 error) *AuthService_ListAccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Login provides a mock function with given fields: ctx, email, password
func (_m *AuthService) Login(ctx context.Context, email string, password string) (auth.Auth, bool, error) {
	ret := _m.Called(ctx, email, password)

	var r0 auth.Auth
	if rf, ok := ret.Get(0).(func(context.Context, string, string) auth.Auth); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(auth.Auth)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, string, string) bool); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AuthService_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type AuthService_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
//   - password string
func (_e *AuthService_Expecter) Login(ctx interface{}, email interface{}, password interface{}) *AuthService_Login_Call {
	return &AuthService_Login_Call{Call: _e.mock.On("Login", ctx, email, password)}
}

func (_c *AuthService_Login_Call) Run(run func(ctx context.Context, email string, password string)) *AuthService_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AuthService_Login_Call) Return(_a0 auth.Auth, _a1 bool, _a2 error) *AuthService_Login_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

type mockConstructorTestingTNewAuthService interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthService(t mockConstructorTestingTNewAuthService) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
