// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/debit"
	mock "github.com/stretchr/testify/mock"
)

// NewDebitUseCase creates a new instance of DebitUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDebitUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *DebitUseCase {
	mock := &DebitUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// DebitUseCase is an autogenerated mock type for the DebitUseCase type
type DebitUseCase struct {
	mock.Mock
}

type DebitUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *DebitUseCase) EXPECT() *DebitUseCase_Expecter {
	return &DebitUseCase_Expecter{mock: &_m.Mock}
}

// GetDebitCardsByUserID provides a mock function for the type DebitUseCase
func (_mock *DebitUseCase) GetDebitCardsByUserID(userID string, limit int, offset int) ([]*debit.DebitCardResponse, int, error) {
	ret := _mock.Called(userID, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetDebitCardsByUserID")
	}

	var r0 []*debit.DebitCardResponse
	var r1 int
	var r2 error
	if returnFunc, ok := ret.Get(0).(func(string, int, int) ([]*debit.DebitCardResponse, int, error)); ok {
		return returnFunc(userID, limit, offset)
	}
	if returnFunc, ok := ret.Get(0).(func(string, int, int) []*debit.DebitCardResponse); ok {
		r0 = returnFunc(userID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*debit.DebitCardResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(string, int, int) int); ok {
		r1 = returnFunc(userID, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}
	if returnFunc, ok := ret.Get(2).(func(string, int, int) error); ok {
		r2 = returnFunc(userID, limit, offset)
	} else {
		r2 = ret.Error(2)
	}
	return r0, r1, r2
}

// DebitUseCase_GetDebitCardsByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDebitCardsByUserID'
type DebitUseCase_GetDebitCardsByUserID_Call struct {
	*mock.Call
}

// GetDebitCardsByUserID is a helper method to define mock.On call
//   - userID string
//   - limit int
//   - offset int
func (_e *DebitUseCase_Expecter) GetDebitCardsByUserID(userID interface{}, limit interface{}, offset interface{}) *DebitUseCase_GetDebitCardsByUserID_Call {
	return &DebitUseCase_GetDebitCardsByUserID_Call{Call: _e.mock.On("GetDebitCardsByUserID", userID, limit, offset)}
}

func (_c *DebitUseCase_GetDebitCardsByUserID_Call) Run(run func(userID string, limit int, offset int)) *DebitUseCase_GetDebitCardsByUserID_Call {
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

func (_c *DebitUseCase_GetDebitCardsByUserID_Call) Return(debitCardResponses []*debit.DebitCardResponse, n int, err error) *DebitUseCase_GetDebitCardsByUserID_Call {
	_c.Call.Return(debitCardResponses, n, err)
	return _c
}

func (_c *DebitUseCase_GetDebitCardsByUserID_Call) RunAndReturn(run func(userID string, limit int, offset int) ([]*debit.DebitCardResponse, int, error)) *DebitUseCase_GetDebitCardsByUserID_Call {
	_c.Call.Return(run)
	return _c
}
