// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package connection

import (
	"context"
	"github.com/apicurio/apicurio-cli/pkg/shared/connection/api"
	"sync"
)

// Ensure, that ConnectionMock does implement Connection.
// If this is not the case, regenerate this file with moq.
var _ Connection = &ConnectionMock{}

// ConnectionMock is a mock implementation of Connection.
//
//	func TestSomethingThatUsesConnection(t *testing.T) {
//
//		// make and configure a mocked Connection
//		mockedConnection := &ConnectionMock{
//			APIFunc: func() apis.API {
//				panic("mock out the API method")
//			},
//			LogoutFunc: func(ctx context.Context) error {
//				panic("mock out the Logout method")
//			},
//			RefreshTokensFunc: func(ctx context.Context) error {
//				panic("mock out the RefreshTokens method")
//			},
//		}
//
//		// use mockedConnection in code that requires Connection
//		// and then make assertions.
//
//	}
type ConnectionMock struct {
	// APIFunc mocks the API method.
	APIFunc func() api.API

	// LogoutFunc mocks the Logout method.
	LogoutFunc func(ctx context.Context) error

	// RefreshTokensFunc mocks the RefreshTokens method.
	RefreshTokensFunc func(ctx context.Context) error

	// calls tracks calls to the methods.
	calls struct {
		// API holds details about calls to the API method.
		API []struct {
		}
		// Logout holds details about calls to the Logout method.
		Logout []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// RefreshTokens holds details about calls to the RefreshTokens method.
		RefreshTokens []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockAPI           sync.RWMutex
	lockLogout        sync.RWMutex
	lockRefreshTokens sync.RWMutex
}

// API calls APIFunc.
func (mock *ConnectionMock) API() api.API {
	if mock.APIFunc == nil {
		panic("ConnectionMock.APIFunc: method is nil but Connection.API was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAPI.Lock()
	mock.calls.API = append(mock.calls.API, callInfo)
	mock.lockAPI.Unlock()
	return mock.APIFunc()
}

// APICalls gets all the calls that were made to API.
// Check the length with:
//
//	len(mockedConnection.APICalls())
func (mock *ConnectionMock) APICalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAPI.RLock()
	calls = mock.calls.API
	mock.lockAPI.RUnlock()
	return calls
}

// Logout calls LogoutFunc.
func (mock *ConnectionMock) Logout(ctx context.Context) error {
	if mock.LogoutFunc == nil {
		panic("ConnectionMock.LogoutFunc: method is nil but Connection.Logout was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockLogout.Lock()
	mock.calls.Logout = append(mock.calls.Logout, callInfo)
	mock.lockLogout.Unlock()
	return mock.LogoutFunc(ctx)
}

// LogoutCalls gets all the calls that were made to Logout.
// Check the length with:
//
//	len(mockedConnection.LogoutCalls())
func (mock *ConnectionMock) LogoutCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockLogout.RLock()
	calls = mock.calls.Logout
	mock.lockLogout.RUnlock()
	return calls
}

// RefreshTokens calls RefreshTokensFunc.
func (mock *ConnectionMock) RefreshTokens(ctx context.Context) error {
	if mock.RefreshTokensFunc == nil {
		panic("ConnectionMock.RefreshTokensFunc: method is nil but Connection.RefreshTokens was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockRefreshTokens.Lock()
	mock.calls.RefreshTokens = append(mock.calls.RefreshTokens, callInfo)
	mock.lockRefreshTokens.Unlock()
	return mock.RefreshTokensFunc(ctx)
}

// RefreshTokensCalls gets all the calls that were made to RefreshTokens.
// Check the length with:
//
//	len(mockedConnection.RefreshTokensCalls())
func (mock *ConnectionMock) RefreshTokensCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockRefreshTokens.RLock()
	calls = mock.calls.RefreshTokens
	mock.lockRefreshTokens.RUnlock()
	return calls
}
