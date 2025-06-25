package errs_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	testCases := []struct {
		name     string
		apiError *errs.APIError
		expected string
	}{
		{
			name:     "SimpleMessage",
			apiError: errs.New(http.StatusBadRequest, "invalid input", nil),
			expected: "invalid input",
		},
		{
			name:     "EmptyMessage",
			apiError: errs.New(http.StatusInternalServerError, "", nil),
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.apiError.Error()
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnwrap(t *testing.T) {
	testCases := []struct {
		name     string
		apiError *errs.APIError
		expected error
	}{
		{
			name:     "WithWrappedError",
			apiError: errs.New(http.StatusInternalServerError, "database error", errors.New("connection failed")),
			expected: errors.New("connection failed"),
		},
		{
			name:     "WithoutWrappedError",
			apiError: errs.New(http.StatusBadRequest, "invalid input", nil),
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.apiError.Unwrap()
			if tc.expected == nil {
				require.Nil(t, result)
			} else {
				require.Equal(t, tc.expected.Error(), result.Error())
			}
		})
	}
}

func TestNew(t *testing.T) {
	testCases := []struct {
		name         string
		code         int
		message      string
		err          error
		expectedCode int
		expectedMsg  string
	}{
		{
			name:         "BasicError",
			code:         http.StatusBadRequest,
			message:      "invalid request",
			err:          errors.New("validation failed"),
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "invalid request",
		},
		{
			name:         "ErrorWithoutWrappedError",
			code:         http.StatusNotFound,
			message:      "resource not found",
			err:          nil,
			expectedCode: http.StatusNotFound,
			expectedMsg:  "resource not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiError := errs.New(tc.code, tc.message, tc.err)

			require.Equal(t, tc.expectedCode, apiError.Code)
			require.Equal(t, tc.expectedMsg, apiError.Message)
			require.Equal(t, tc.err, apiError.Err)
		})
	}
}

func TestNotFound(t *testing.T) {
	message := "user not found"
	apiError := errs.NotFound(message)

	require.Equal(t, http.StatusNotFound, apiError.Code)
	require.Equal(t, message, apiError.Message)
	require.NotNil(t, apiError.Err)
	require.Equal(t, message, apiError.Err.Error())
}

func TestInternal(t *testing.T) {
	originalErr := errors.New("database connection failed")
	apiError := errs.Internal(originalErr)

	require.Equal(t, http.StatusInternalServerError, apiError.Code)
	require.Equal(t, "internal server error", apiError.Message)
	require.Equal(t, originalErr, apiError.Err)
}

func TestBadRequest(t *testing.T) {
	message := "invalid input data"
	apiError := errs.BadRequest(message)

	require.Equal(t, http.StatusBadRequest, apiError.Code)
	require.Equal(t, message, apiError.Message)
	require.NotNil(t, apiError.Err)
	require.Equal(t, message, apiError.Err.Error())
}

func TestUnauthorized(t *testing.T) {
	message := "access denied"
	apiError := errs.Unauthorized(message)

	require.Equal(t, http.StatusUnauthorized, apiError.Code)
	require.Equal(t, message, apiError.Message)
	require.NotNil(t, apiError.Err)
	require.Equal(t, message, apiError.Err.Error())
}
