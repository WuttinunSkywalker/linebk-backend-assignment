package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestErrorHandler(t *testing.T) {
	testCases := []struct {
		name         string
		setupError   func(ctx *gin.Context)
		expectedCode int
		expectedBody string
	}{
		{
			name: "NoError",
			setupError: func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"message": "success"})
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"message":"success"}`,
		},
		{
			name: "APIError",
			setupError: func(ctx *gin.Context) {
				ctx.Error(errs.BadRequest("Invalid request data"))
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"success":false,"message":"Invalid request data"}`,
		},
		{
			name: "UnauthorizedError",
			setupError: func(ctx *gin.Context) {
				ctx.Error(errs.Unauthorized("Access denied"))
			},
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"success":false,"message":"Access denied"}`,
		},
		{
			name: "NotFoundError",
			setupError: func(ctx *gin.Context) {
				ctx.Error(errs.NotFound("Resource not found"))
			},
			expectedCode: http.StatusNotFound,
			expectedBody: `{"success":false,"message":"Resource not found"}`,
		},
		{
			name: "InternalError",
			setupError: func(ctx *gin.Context) {
				ctx.Error(errs.Internal(errors.New("database connection failed")))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"success":false,"message":"internal server error"}`,
		},
		{
			name: "UnexpectedError",
			setupError: func(ctx *gin.Context) {
				ctx.Error(errors.New("unexpected system error"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"success":false,"message":"An unexpected internal server error occurred"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.Use(middleware.ErrorHandler())

			testPath := "/test"
			router.GET(testPath, func(ctx *gin.Context) {
				tc.setupError(ctx)
			})

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, testPath, nil)
			require.NoError(t, err)

			router.ServeHTTP(recorder, request)

			require.Equal(t, tc.expectedCode, recorder.Code)
			require.JSONEq(t, tc.expectedBody, recorder.Body.String())
		})
	}
}
