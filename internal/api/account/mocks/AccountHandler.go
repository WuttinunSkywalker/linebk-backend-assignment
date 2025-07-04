// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// NewAccountHandler creates a new instance of AccountHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountHandler {
	mock := &AccountHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// AccountHandler is an autogenerated mock type for the AccountHandler type
type AccountHandler struct {
	mock.Mock
}

type AccountHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountHandler) EXPECT() *AccountHandler_Expecter {
	return &AccountHandler_Expecter{mock: &_m.Mock}
}

// GetMyAccounts provides a mock function for the type AccountHandler
func (_mock *AccountHandler) GetMyAccounts(ctx *gin.Context) {
	_mock.Called(ctx)
	return
}

// AccountHandler_GetMyAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMyAccounts'
type AccountHandler_GetMyAccounts_Call struct {
	*mock.Call
}

// GetMyAccounts is a helper method to define mock.On call
//   - ctx *gin.Context
func (_e *AccountHandler_Expecter) GetMyAccounts(ctx interface{}) *AccountHandler_GetMyAccounts_Call {
	return &AccountHandler_GetMyAccounts_Call{Call: _e.mock.On("GetMyAccounts", ctx)}
}

func (_c *AccountHandler_GetMyAccounts_Call) Run(run func(ctx *gin.Context)) *AccountHandler_GetMyAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 *gin.Context
		if args[0] != nil {
			arg0 = args[0].(*gin.Context)
		}
		run(
			arg0,
		)
	})
	return _c
}

func (_c *AccountHandler_GetMyAccounts_Call) Return() *AccountHandler_GetMyAccounts_Call {
	_c.Call.Return()
	return _c
}

func (_c *AccountHandler_GetMyAccounts_Call) RunAndReturn(run func(ctx *gin.Context)) *AccountHandler_GetMyAccounts_Call {
	_c.Run(run)
	return _c
}
