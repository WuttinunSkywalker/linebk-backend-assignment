// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/banner"
	mock "github.com/stretchr/testify/mock"
)

// NewBannerRepository creates a new instance of BannerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBannerRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BannerRepository {
	mock := &BannerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// BannerRepository is an autogenerated mock type for the BannerRepository type
type BannerRepository struct {
	mock.Mock
}

type BannerRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *BannerRepository) EXPECT() *BannerRepository_Expecter {
	return &BannerRepository_Expecter{mock: &_m.Mock}
}

// CountBannersByUserID provides a mock function for the type BannerRepository
func (_mock *BannerRepository) CountBannersByUserID(userID string) (int, error) {
	ret := _mock.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for CountBannersByUserID")
	}

	var r0 int
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(string) (int, error)); ok {
		return returnFunc(userID)
	}
	if returnFunc, ok := ret.Get(0).(func(string) int); ok {
		r0 = returnFunc(userID)
	} else {
		r0 = ret.Get(0).(int)
	}
	if returnFunc, ok := ret.Get(1).(func(string) error); ok {
		r1 = returnFunc(userID)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// BannerRepository_CountBannersByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountBannersByUserID'
type BannerRepository_CountBannersByUserID_Call struct {
	*mock.Call
}

// CountBannersByUserID is a helper method to define mock.On call
//   - userID string
func (_e *BannerRepository_Expecter) CountBannersByUserID(userID interface{}) *BannerRepository_CountBannersByUserID_Call {
	return &BannerRepository_CountBannersByUserID_Call{Call: _e.mock.On("CountBannersByUserID", userID)}
}

func (_c *BannerRepository_CountBannersByUserID_Call) Run(run func(userID string)) *BannerRepository_CountBannersByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 string
		if args[0] != nil {
			arg0 = args[0].(string)
		}
		run(
			arg0,
		)
	})
	return _c
}

func (_c *BannerRepository_CountBannersByUserID_Call) Return(n int, err error) *BannerRepository_CountBannersByUserID_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *BannerRepository_CountBannersByUserID_Call) RunAndReturn(run func(userID string) (int, error)) *BannerRepository_CountBannersByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// GetBannersByUserID provides a mock function for the type BannerRepository
func (_mock *BannerRepository) GetBannersByUserID(userID string, limit int, offset int) ([]*banner.Banner, error) {
	ret := _mock.Called(userID, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetBannersByUserID")
	}

	var r0 []*banner.Banner
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(string, int, int) ([]*banner.Banner, error)); ok {
		return returnFunc(userID, limit, offset)
	}
	if returnFunc, ok := ret.Get(0).(func(string, int, int) []*banner.Banner); ok {
		r0 = returnFunc(userID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*banner.Banner)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = returnFunc(userID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// BannerRepository_GetBannersByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBannersByUserID'
type BannerRepository_GetBannersByUserID_Call struct {
	*mock.Call
}

// GetBannersByUserID is a helper method to define mock.On call
//   - userID string
//   - limit int
//   - offset int
func (_e *BannerRepository_Expecter) GetBannersByUserID(userID interface{}, limit interface{}, offset interface{}) *BannerRepository_GetBannersByUserID_Call {
	return &BannerRepository_GetBannersByUserID_Call{Call: _e.mock.On("GetBannersByUserID", userID, limit, offset)}
}

func (_c *BannerRepository_GetBannersByUserID_Call) Run(run func(userID string, limit int, offset int)) *BannerRepository_GetBannersByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 string
		if args[0] != nil {
			arg0 = args[0].(string)
		}
		var arg1 int
		if args[1] != nil {
			arg1 = args[1].(int)
		}
		var arg2 int
		if args[2] != nil {
			arg2 = args[2].(int)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *BannerRepository_GetBannersByUserID_Call) Return(banners []*banner.Banner, err error) *BannerRepository_GetBannersByUserID_Call {
	_c.Call.Return(banners, err)
	return _c
}

func (_c *BannerRepository_GetBannersByUserID_Call) RunAndReturn(run func(userID string, limit int, offset int) ([]*banner.Banner, error)) *BannerRepository_GetBannersByUserID_Call {
	_c.Call.Return(run)
	return _c
}
