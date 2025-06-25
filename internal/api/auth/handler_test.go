package auth_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/auth"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/auth/mocks"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAuthHandler_AuthHandlerLoginWithPin(t *testing.T) {
	testLoginResponse := &auth.LoginResponse{
		AccessToken:  "test-access-token",
		RefreshToken: "test-refresh-token",
	}

	testCases := []struct {
		name          string
		requestBody   string
		buildStubs    func(usecase *mocks.AuthUsecase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			requestBody: `{"user_id":"test-user-123","pin":"123456"}`,
			buildStubs: func(usecase *mocks.AuthUsecase) {
				usecase.EXPECT().LoginWithPin("test-user-123", "123456").Return(testLoginResponse, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:        "BadRequest_InvalidJSON",
			requestBody: `{"user_id":"test-user-123","pin":}`,
			buildStubs: func(usecase *mocks.AuthUsecase) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:        "BadRequest_MissingFields",
			requestBody: `{"user_id":"test-user-123"}`,
			buildStubs: func(usecase *mocks.AuthUsecase) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:        "Unauthorized_InvalidPin",
			requestBody: `{"user_id":"test-user-123","pin":"123456"}`,
			buildStubs: func(usecase *mocks.AuthUsecase) {
				usecase.EXPECT().LoginWithPin("test-user-123", "123456").Return(nil, errs.Unauthorized("invalid pin"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:        "InternalServerError",
			requestBody: `{"user_id":"test-user-123","pin":"123456"}`,
			buildStubs: func(usecase *mocks.AuthUsecase) {
				usecase.EXPECT().LoginWithPin("test-user-123", "123456").Return(nil, errs.Internal(errors.New("database error")))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewAuthUsecase(t)
			tc.buildStubs(mockUsecase)

			handler := auth.NewAuthHandler(mockUsecase)

			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			engine := gin.New()
			engine.Use(middleware.ErrorHandler())

			engine.POST("/auth/login", handler.LoginWithPin)

			req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString(tc.requestBody))
			req.Header.Set("Content-Type", "application/json")

			engine.ServeHTTP(recorder, req)

			tc.checkResponse(t, recorder)
		})
	}
}
