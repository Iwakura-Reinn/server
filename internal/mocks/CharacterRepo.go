// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	character "github.com/bangumi/server/internal/character"

	domain "github.com/bangumi/server/domain"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"
)

// CharacterRepo is an autogenerated mock type for the Repo type
type CharacterRepo struct {
	mock.Mock
}

type CharacterRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *CharacterRepo) EXPECT() *CharacterRepo_Expecter {
	return &CharacterRepo_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *CharacterRepo) Get(ctx context.Context, id uint32) (model.Character, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Character
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) (model.Character, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.Character); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Character)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CharacterRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CharacterRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
func (_e *CharacterRepo_Expecter) Get(ctx interface{}, id interface{}) *CharacterRepo_Get_Call {
	return &CharacterRepo_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *CharacterRepo_Get_Call) Run(run func(ctx context.Context, id uint32)) *CharacterRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *CharacterRepo_Get_Call) Return(_a0 model.Character, _a1 error) *CharacterRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CharacterRepo_Get_Call) RunAndReturn(run func(context.Context, uint32) (model.Character, error)) *CharacterRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByIDs provides a mock function with given fields: ctx, ids
func (_m *CharacterRepo) GetByIDs(ctx context.Context, ids []uint32) (map[uint32]model.Character, error) {
	ret := _m.Called(ctx, ids)

	var r0 map[uint32]model.Character
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32) (map[uint32]model.Character, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint32) map[uint32]model.Character); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]model.Character)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint32) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CharacterRepo_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type CharacterRepo_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []uint32
func (_e *CharacterRepo_Expecter) GetByIDs(ctx interface{}, ids interface{}) *CharacterRepo_GetByIDs_Call {
	return &CharacterRepo_GetByIDs_Call{Call: _e.mock.On("GetByIDs", ctx, ids)}
}

func (_c *CharacterRepo_GetByIDs_Call) Run(run func(ctx context.Context, ids []uint32)) *CharacterRepo_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uint32))
	})
	return _c
}

func (_c *CharacterRepo_GetByIDs_Call) Return(_a0 map[uint32]model.Character, _a1 error) *CharacterRepo_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CharacterRepo_GetByIDs_Call) RunAndReturn(run func(context.Context, []uint32) (map[uint32]model.Character, error)) *CharacterRepo_GetByIDs_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonRelated provides a mock function with given fields: ctx, personID
func (_m *CharacterRepo) GetPersonRelated(ctx context.Context, personID uint32) ([]domain.PersonCharacterRelation, error) {
	ret := _m.Called(ctx, personID)

	var r0 []domain.PersonCharacterRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) ([]domain.PersonCharacterRelation, error)); ok {
		return rf(ctx, personID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []domain.PersonCharacterRelation); ok {
		r0 = rf(ctx, personID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PersonCharacterRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, personID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CharacterRepo_GetPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonRelated'
type CharacterRepo_GetPersonRelated_Call struct {
	*mock.Call
}

// GetPersonRelated is a helper method to define mock.On call
//   - ctx context.Context
//   - personID uint32
func (_e *CharacterRepo_Expecter) GetPersonRelated(ctx interface{}, personID interface{}) *CharacterRepo_GetPersonRelated_Call {
	return &CharacterRepo_GetPersonRelated_Call{Call: _e.mock.On("GetPersonRelated", ctx, personID)}
}

func (_c *CharacterRepo_GetPersonRelated_Call) Run(run func(ctx context.Context, personID uint32)) *CharacterRepo_GetPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *CharacterRepo_GetPersonRelated_Call) Return(_a0 []domain.PersonCharacterRelation, _a1 error) *CharacterRepo_GetPersonRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CharacterRepo_GetPersonRelated_Call) RunAndReturn(run func(context.Context, uint32) ([]domain.PersonCharacterRelation, error)) *CharacterRepo_GetPersonRelated_Call {
	_c.Call.Return(run)
	return _c
}

// GetRelations provides a mock function with given fields: ctx, ids
func (_m *CharacterRepo) GetRelations(ctx context.Context, ids []character.CompositeId) ([]domain.SubjectCharacterRelation, error) {
	ret := _m.Called(ctx, ids)

	var r0 []domain.SubjectCharacterRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []character.CompositeId) ([]domain.SubjectCharacterRelation, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []character.CompositeId) []domain.SubjectCharacterRelation); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SubjectCharacterRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []character.CompositeId) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CharacterRepo_GetRelations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRelations'
type CharacterRepo_GetRelations_Call struct {
	*mock.Call
}

// GetRelations is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []character.CompositeId
func (_e *CharacterRepo_Expecter) GetRelations(ctx interface{}, ids interface{}) *CharacterRepo_GetRelations_Call {
	return &CharacterRepo_GetRelations_Call{Call: _e.mock.On("GetRelations", ctx, ids)}
}

func (_c *CharacterRepo_GetRelations_Call) Run(run func(ctx context.Context, ids []character.CompositeId)) *CharacterRepo_GetRelations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]character.CompositeId))
	})
	return _c
}

func (_c *CharacterRepo_GetRelations_Call) Return(_a0 []domain.SubjectCharacterRelation, _a1 error) *CharacterRepo_GetRelations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CharacterRepo_GetRelations_Call) RunAndReturn(run func(context.Context, []character.CompositeId) ([]domain.SubjectCharacterRelation, error)) *CharacterRepo_GetRelations_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubjectRelated provides a mock function with given fields: ctx, subjectID
func (_m *CharacterRepo) GetSubjectRelated(ctx context.Context, subjectID uint32) ([]domain.SubjectCharacterRelation, error) {
	ret := _m.Called(ctx, subjectID)

	var r0 []domain.SubjectCharacterRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) ([]domain.SubjectCharacterRelation, error)); ok {
		return rf(ctx, subjectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []domain.SubjectCharacterRelation); ok {
		r0 = rf(ctx, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SubjectCharacterRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CharacterRepo_GetSubjectRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectRelated'
type CharacterRepo_GetSubjectRelated_Call struct {
	*mock.Call
}

// GetSubjectRelated is a helper method to define mock.On call
//   - ctx context.Context
//   - subjectID uint32
func (_e *CharacterRepo_Expecter) GetSubjectRelated(ctx interface{}, subjectID interface{}) *CharacterRepo_GetSubjectRelated_Call {
	return &CharacterRepo_GetSubjectRelated_Call{Call: _e.mock.On("GetSubjectRelated", ctx, subjectID)}
}

func (_c *CharacterRepo_GetSubjectRelated_Call) Run(run func(ctx context.Context, subjectID uint32)) *CharacterRepo_GetSubjectRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *CharacterRepo_GetSubjectRelated_Call) Return(_a0 []domain.SubjectCharacterRelation, _a1 error) *CharacterRepo_GetSubjectRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CharacterRepo_GetSubjectRelated_Call) RunAndReturn(run func(context.Context, uint32) ([]domain.SubjectCharacterRelation, error)) *CharacterRepo_GetSubjectRelated_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCharacterRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewCharacterRepo creates a new instance of CharacterRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCharacterRepo(t mockConstructorTestingTNewCharacterRepo) *CharacterRepo {
	mock := &CharacterRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
