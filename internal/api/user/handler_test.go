package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/user"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/user/mocks"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const testUserID = "test-user-123"

func TestUserHandler_GetMe(t *testing.T) {
	testTime := time.Now()
	testImage := "test-image.jpg"

	testUserResponse := &user.UserResponse{
		UserID:    testUserID,
		Name:      "Test User",
		Image:     &testImage,
		CreatedAt: testTime,
		UpdatedAt: testTime,
	}

	testCases := []struct {
		name          string
		setupContext  func(ctx *gin.Context)
		buildStubs    func(usecase *mocks.UserUsecase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupContext: func(ctx *gin.Context) {
				ctx.Set("userID", testUserID)
			},
			buildStubs: func(usecase *mocks.UserUsecase) {
				usecase.EXPECT().GetUserByID(testUserID).Return(testUserResponse, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Unauthorized",
			setupContext: func(ctx *gin.Context) {
			},
			buildStubs: func(usecase *mocks.UserUsecase) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			setupContext: func(ctx *gin.Context) {
				ctx.Set("userID", testUserID)
			},
			buildStubs: func(usecase *mocks.UserUsecase) {
				usecase.EXPECT().GetUserByID(testUserID).Return(nil, errors.New("database error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewUserUsecase(t)
			tc.buildStubs(mockUsecase)

			handler := user.NewUserHandler(mockUsecase)

			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			engine := gin.New()
			engine.Use(middleware.ErrorHandler())

			engine.GET("/me", func(ctx *gin.Context) {
				tc.setupContext(ctx)
				handler.GetMe(ctx)
			})

			req, _ := http.NewRequest(http.MethodGet, "/me", nil)
			engine.ServeHTTP(recorder, req)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestUserHandler_GetMyGreeting(t *testing.T) {
	testTime := time.Now()

	testGreetingResponse := &user.UserGreetingResponse{
		UserID:    testUserID,
		Greeting:  "Hello, World!",
		CreatedAt: testTime,
	}

	testCases := []struct {
		name          string
		setupContext  func(ctx *gin.Context)
		buildStubs    func(usecase *mocks.UserUsecase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupContext: func(ctx *gin.Context) {
				ctx.Set("userID", testUserID)
			},
			buildStubs: func(usecase *mocks.UserUsecase) {
				usecase.EXPECT().GetUserGreetingByUserID(testUserID).Return(testGreetingResponse, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Unauthorized",
			setupContext: func(ctx *gin.Context) {
			},
			buildStubs: func(usecase *mocks.UserUsecase) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			setupContext: func(ctx *gin.Context) {
				ctx.Set("userID", testUserID)
			},
			buildStubs: func(usecase *mocks.UserUsecase) {
				usecase.EXPECT().GetUserGreetingByUserID(testUserID).Return(nil, errors.New("database error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewUserUsecase(t)
			tc.buildStubs(mockUsecase)

			handler := user.NewUserHandler(mockUsecase)

			gin.SetMode(gin.TestMode)
			recorder := httptest.NewRecorder()
			engine := gin.New()
			engine.Use(middleware.ErrorHandler())

			engine.GET("/me/greeting", func(ctx *gin.Context) {
				tc.setupContext(ctx)
				handler.GetMyGreeting(ctx)
			})

			req, _ := http.NewRequest(http.MethodGet, "/me/greeting", nil)
			engine.ServeHTTP(recorder, req)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestUserHandler_GetUserPreview(t *testing.T) {
	testImage := "test-preview-image.jpg"
	testUserPreviewResponse := &user.UserPreviewResponse{
		Name:  "Test Preview User",
		Image: &testImage,
	}

	testCases := []struct {
		name          string
		userID        string
		buildStubs    func(usecase *mocks.UserUsecase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userID: testUserID,
			buildStubs: func(usecase *mocks.UserUsecase) {
				usecase.EXPECT().GetUserPreview(testUserID).Return(testUserPreviewResponse, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:       "BadRequest_EmptyUserID",
			userID:     "",
			buildStubs: func(usecase *mocks.UserUsecase) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "NotFound",
			userID: "non-existent-user",
			buildStubs: func(usecase *mocks.UserUsecase) {
				usecase.EXPECT().GetUserPreview("non-existent-user").Return(nil, errors.New("user not found"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "InternalServerError",
			userID: testUserID,
			buildStubs: func(usecase *mocks.UserUsecase) {
				usecase.EXPECT().GetUserPreview(testUserID).Return(nil, errors.New("database error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewUserUsecase(t)
			tc.buildStubs(mockUsecase)

			handler := user.NewUserHandler(mockUsecase)

			gin.SetMode(gin.TestMode)
			recorder := httptest.NewRecorder()
			engine := gin.New()
			engine.Use(middleware.ErrorHandler())

			engine.GET("/users/:userid/preview", handler.GetUserPreview)

			req, _ := http.NewRequest(http.MethodGet, "/users/"+tc.userID+"/preview", nil)
			engine.ServeHTTP(recorder, req)

			tc.checkResponse(t, recorder)
		})
	}
}
