// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/internal/domain"
	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"

	query "github.com/bangumi/server/internal/dal/query"

	time "time"
)

// CollectionRepo is an autogenerated mock type for the CollectionRepo type
type CollectionRepo struct {
	mock.Mock
}

type CollectionRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *CollectionRepo) EXPECT() *CollectionRepo_Expecter {
	return &CollectionRepo_Expecter{mock: &_m.Mock}
}

// CountSubjectCollections provides a mock function with given fields: ctx, userID, subjectType, collectionType, showPrivate
func (_m *CollectionRepo) CountSubjectCollections(ctx context.Context, userID model.UserID, subjectType uint8, collectionType model.SubjectCollection, showPrivate bool) (int64, error) {
	ret := _m.Called(ctx, userID, subjectType, collectionType, showPrivate)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, uint8, model.SubjectCollection, bool) int64); ok {
		r0 = rf(ctx, userID, subjectType, collectionType, showPrivate)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, uint8, model.SubjectCollection, bool) error); ok {
		r1 = rf(ctx, userID, subjectType, collectionType, showPrivate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_CountSubjectCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountSubjectCollections'
type CollectionRepo_CountSubjectCollections_Call struct {
	*mock.Call
}

// CountSubjectCollections is a helper method to define mock.On call
//   - ctx context.Context
//   - userID model.UserID
//   - subjectType uint8
//   - collectionType model.SubjectCollection
//   - showPrivate bool
func (_e *CollectionRepo_Expecter) CountSubjectCollections(ctx interface{}, userID interface{}, subjectType interface{}, collectionType interface{}, showPrivate interface{}) *CollectionRepo_CountSubjectCollections_Call {
	return &CollectionRepo_CountSubjectCollections_Call{Call: _e.mock.On("CountSubjectCollections", ctx, userID, subjectType, collectionType, showPrivate)}
}

func (_c *CollectionRepo_CountSubjectCollections_Call) Run(run func(ctx context.Context, userID model.UserID, subjectType uint8, collectionType model.SubjectCollection, showPrivate bool)) *CollectionRepo_CountSubjectCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(uint8), args[3].(model.SubjectCollection), args[4].(bool))
	})
	return _c
}

func (_c *CollectionRepo_CountSubjectCollections_Call) Return(_a0 int64, _a1 error) *CollectionRepo_CountSubjectCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSubjectCollection provides a mock function with given fields: ctx, userID, subjectID
func (_m *CollectionRepo) GetSubjectCollection(ctx context.Context, userID model.UserID, subjectID model.SubjectID) (model.UserSubjectCollection, error) {
	ret := _m.Called(ctx, userID, subjectID)

	var r0 model.UserSubjectCollection
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, model.SubjectID) model.UserSubjectCollection); ok {
		r0 = rf(ctx, userID, subjectID)
	} else {
		r0 = ret.Get(0).(model.UserSubjectCollection)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, model.SubjectID) error); ok {
		r1 = rf(ctx, userID, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_GetSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectCollection'
type CollectionRepo_GetSubjectCollection_Call struct {
	*mock.Call
}

// GetSubjectCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID model.UserID
//   - subjectID model.SubjectID
func (_e *CollectionRepo_Expecter) GetSubjectCollection(ctx interface{}, userID interface{}, subjectID interface{}) *CollectionRepo_GetSubjectCollection_Call {
	return &CollectionRepo_GetSubjectCollection_Call{Call: _e.mock.On("GetSubjectCollection", ctx, userID, subjectID)}
}

func (_c *CollectionRepo_GetSubjectCollection_Call) Run(run func(ctx context.Context, userID model.UserID, subjectID model.SubjectID)) *CollectionRepo_GetSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(model.SubjectID))
	})
	return _c
}

func (_c *CollectionRepo_GetSubjectCollection_Call) Return(_a0 model.UserSubjectCollection, _a1 error) *CollectionRepo_GetSubjectCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSubjectEpisodesCollection provides a mock function with given fields: ctx, userID, subjectID
func (_m *CollectionRepo) GetSubjectEpisodesCollection(ctx context.Context, userID model.UserID, subjectID model.SubjectID) (model.UserSubjectEpisodesCollection, error) {
	ret := _m.Called(ctx, userID, subjectID)

	var r0 model.UserSubjectEpisodesCollection
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, model.SubjectID) model.UserSubjectEpisodesCollection); ok {
		r0 = rf(ctx, userID, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.UserSubjectEpisodesCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, model.SubjectID) error); ok {
		r1 = rf(ctx, userID, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_GetSubjectEpisodesCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectEpisodesCollection'
type CollectionRepo_GetSubjectEpisodesCollection_Call struct {
	*mock.Call
}

// GetSubjectEpisodesCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID model.UserID
//   - subjectID model.SubjectID
func (_e *CollectionRepo_Expecter) GetSubjectEpisodesCollection(ctx interface{}, userID interface{}, subjectID interface{}) *CollectionRepo_GetSubjectEpisodesCollection_Call {
	return &CollectionRepo_GetSubjectEpisodesCollection_Call{Call: _e.mock.On("GetSubjectEpisodesCollection", ctx, userID, subjectID)}
}

func (_c *CollectionRepo_GetSubjectEpisodesCollection_Call) Run(run func(ctx context.Context, userID model.UserID, subjectID model.SubjectID)) *CollectionRepo_GetSubjectEpisodesCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(model.SubjectID))
	})
	return _c
}

func (_c *CollectionRepo_GetSubjectEpisodesCollection_Call) Return(_a0 model.UserSubjectEpisodesCollection, _a1 error) *CollectionRepo_GetSubjectEpisodesCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListSubjectCollection provides a mock function with given fields: ctx, userID, subjectType, collectionType, showPrivate, limit, offset
func (_m *CollectionRepo) ListSubjectCollection(ctx context.Context, userID model.UserID, subjectType uint8, collectionType model.SubjectCollection, showPrivate bool, limit int, offset int) ([]model.UserSubjectCollection, error) {
	ret := _m.Called(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)

	var r0 []model.UserSubjectCollection
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, uint8, model.SubjectCollection, bool, int, int) []model.UserSubjectCollection); ok {
		r0 = rf(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.UserSubjectCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, uint8, model.SubjectCollection, bool, int, int) error); ok {
		r1 = rf(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_ListSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSubjectCollection'
type CollectionRepo_ListSubjectCollection_Call struct {
	*mock.Call
}

// ListSubjectCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID model.UserID
//   - subjectType uint8
//   - collectionType model.SubjectCollection
//   - showPrivate bool
//   - limit int
//   - offset int
func (_e *CollectionRepo_Expecter) ListSubjectCollection(ctx interface{}, userID interface{}, subjectType interface{}, collectionType interface{}, showPrivate interface{}, limit interface{}, offset interface{}) *CollectionRepo_ListSubjectCollection_Call {
	return &CollectionRepo_ListSubjectCollection_Call{Call: _e.mock.On("ListSubjectCollection", ctx, userID, subjectType, collectionType, showPrivate, limit, offset)}
}

func (_c *CollectionRepo_ListSubjectCollection_Call) Run(run func(ctx context.Context, userID model.UserID, subjectType uint8, collectionType model.SubjectCollection, showPrivate bool, limit int, offset int)) *CollectionRepo_ListSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(uint8), args[3].(model.SubjectCollection), args[4].(bool), args[5].(int), args[6].(int))
	})
	return _c
}

func (_c *CollectionRepo_ListSubjectCollection_Call) Return(_a0 []model.UserSubjectCollection, _a1 error) *CollectionRepo_ListSubjectCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateEpisodeCollection provides a mock function with given fields: ctx, userID, subjectID, episodeIDs, collection, at
func (_m *CollectionRepo) UpdateEpisodeCollection(ctx context.Context, userID model.UserID, subjectID model.SubjectID, episodeIDs []model.EpisodeID, collection model.EpisodeCollection, at time.Time) (model.UserSubjectEpisodesCollection, error) {
	ret := _m.Called(ctx, userID, subjectID, episodeIDs, collection, at)

	var r0 model.UserSubjectEpisodesCollection
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, model.SubjectID, []model.EpisodeID, model.EpisodeCollection, time.Time) model.UserSubjectEpisodesCollection); ok {
		r0 = rf(ctx, userID, subjectID, episodeIDs, collection, at)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.UserSubjectEpisodesCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, model.SubjectID, []model.EpisodeID, model.EpisodeCollection, time.Time) error); ok {
		r1 = rf(ctx, userID, subjectID, episodeIDs, collection, at)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_UpdateEpisodeCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEpisodeCollection'
type CollectionRepo_UpdateEpisodeCollection_Call struct {
	*mock.Call
}

// UpdateEpisodeCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID model.UserID
//   - subjectID model.SubjectID
//   - episodeIDs []model.EpisodeID
//   - collection model.EpisodeCollection
//   - at time.Time
func (_e *CollectionRepo_Expecter) UpdateEpisodeCollection(ctx interface{}, userID interface{}, subjectID interface{}, episodeIDs interface{}, collection interface{}, at interface{}) *CollectionRepo_UpdateEpisodeCollection_Call {
	return &CollectionRepo_UpdateEpisodeCollection_Call{Call: _e.mock.On("UpdateEpisodeCollection", ctx, userID, subjectID, episodeIDs, collection, at)}
}

func (_c *CollectionRepo_UpdateEpisodeCollection_Call) Run(run func(ctx context.Context, userID model.UserID, subjectID model.SubjectID, episodeIDs []model.EpisodeID, collection model.EpisodeCollection, at time.Time)) *CollectionRepo_UpdateEpisodeCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(model.SubjectID), args[3].([]model.EpisodeID), args[4].(model.EpisodeCollection), args[5].(time.Time))
	})
	return _c
}

func (_c *CollectionRepo_UpdateEpisodeCollection_Call) Return(_a0 model.UserSubjectEpisodesCollection, _a1 error) *CollectionRepo_UpdateEpisodeCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateSubjectCollection provides a mock function with given fields: ctx, userID, subjectID, data, at
func (_m *CollectionRepo) UpdateSubjectCollection(ctx context.Context, userID model.UserID, subjectID model.SubjectID, data domain.SubjectCollectionUpdate, at time.Time) error {
	ret := _m.Called(ctx, userID, subjectID, data, at)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, model.SubjectID, domain.SubjectCollectionUpdate, time.Time) error); ok {
		r0 = rf(ctx, userID, subjectID, data, at)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CollectionRepo_UpdateSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSubjectCollection'
type CollectionRepo_UpdateSubjectCollection_Call struct {
	*mock.Call
}

// UpdateSubjectCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID model.UserID
//   - subjectID model.SubjectID
//   - data domain.SubjectCollectionUpdate
//   - at time.Time
func (_e *CollectionRepo_Expecter) UpdateSubjectCollection(ctx interface{}, userID interface{}, subjectID interface{}, data interface{}, at interface{}) *CollectionRepo_UpdateSubjectCollection_Call {
	return &CollectionRepo_UpdateSubjectCollection_Call{Call: _e.mock.On("UpdateSubjectCollection", ctx, userID, subjectID, data, at)}
}

func (_c *CollectionRepo_UpdateSubjectCollection_Call) Run(run func(ctx context.Context, userID model.UserID, subjectID model.SubjectID, data domain.SubjectCollectionUpdate, at time.Time)) *CollectionRepo_UpdateSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(model.SubjectID), args[3].(domain.SubjectCollectionUpdate), args[4].(time.Time))
	})
	return _c
}

func (_c *CollectionRepo_UpdateSubjectCollection_Call) Return(_a0 error) *CollectionRepo_UpdateSubjectCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

// WithQuery provides a mock function with given fields: _a0
func (_m *CollectionRepo) WithQuery(_a0 *query.Query) domain.CollectionRepo {
	ret := _m.Called(_a0)

	var r0 domain.CollectionRepo
	if rf, ok := ret.Get(0).(func(*query.Query) domain.CollectionRepo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.CollectionRepo)
		}
	}

	return r0
}

// CollectionRepo_WithQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithQuery'
type CollectionRepo_WithQuery_Call struct {
	*mock.Call
}

// WithQuery is a helper method to define mock.On call
//   - _a0 *query.Query
func (_e *CollectionRepo_Expecter) WithQuery(_a0 interface{}) *CollectionRepo_WithQuery_Call {
	return &CollectionRepo_WithQuery_Call{Call: _e.mock.On("WithQuery", _a0)}
}

func (_c *CollectionRepo_WithQuery_Call) Run(run func(_a0 *query.Query)) *CollectionRepo_WithQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*query.Query))
	})
	return _c
}

func (_c *CollectionRepo_WithQuery_Call) Return(_a0 domain.CollectionRepo) *CollectionRepo_WithQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewCollectionRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewCollectionRepo creates a new instance of CollectionRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCollectionRepo(t mockConstructorTestingTNewCollectionRepo) *CollectionRepo {
	mock := &CollectionRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
