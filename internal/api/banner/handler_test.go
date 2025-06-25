package banner_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/banner"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/banner/mocks"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const testUserID = "test-user-123"

func TestBannerHandler_GetMyBanners(t *testing.T) {
	testBanners := []*banner.BannerResponse{
		{
			BannerID:    "banner-1",
			UserID:      testUserID,
			Title:       "Test Banner 1",
			Description: "Test Description 1",
		},
		{
			BannerID:    "banner-2",
			UserID:      testUserID,
			Title:       "Test Banner 2",
			Description: "Test Description 2",
		},
	}

	testCases := []struct {
		name          string
		queryParams   string
		setupContext  func(ctx *gin.Context)
		buildStubs    func(usecase *mocks.BannerUsecase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			queryParams: "",
			setupContext: func(ctx *gin.Context) {
				ctx.Set("userID", testUserID)
			},
			buildStubs: func(usecase *mocks.BannerUsecase) {
				usecase.EXPECT().GetBannersByUserID(testUserID, 10, 0).Return(testBanners, 2, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:        "Unauthorized",
			queryParams: "",
			setupContext: func(ctx *gin.Context) {
			},
			buildStubs: func(usecase *mocks.BannerUsecase) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:        "BadRequest_InvalidQuery",
			queryParams: "?limit=invalid",
			setupContext: func(ctx *gin.Context) {
				ctx.Set("userID", testUserID)
			},
			buildStubs: func(usecase *mocks.BannerUsecase) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:        "InternalServerError",
			queryParams: "",
			setupContext: func(ctx *gin.Context) {
				ctx.Set("userID", testUserID)
			},
			buildStubs: func(usecase *mocks.BannerUsecase) {
				usecase.EXPECT().GetBannersByUserID(testUserID, 10, 0).Return(nil, 0, errors.New("database error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewBannerUsecase(t)
			tc.buildStubs(mockUsecase)

			handler := banner.NewBannerHandler(mockUsecase)

			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			engine := gin.New()
			engine.Use(middleware.ErrorHandler())

			engine.GET("/me/banners", func(ctx *gin.Context) {
				tc.setupContext(ctx)
				handler.GetMyBanners(ctx)
			})

			url := "/me/banners" + tc.queryParams
			req, _ := http.NewRequest(http.MethodGet, url, nil)

			engine.ServeHTTP(recorder, req)

			tc.checkResponse(t, recorder)
		})
	}
}
