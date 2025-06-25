package account_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/account"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/account/mocks"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const testUserID = "test-user-123"

func TestAccountHandler_GetMyAccounts(t *testing.T) {
	testAccounts := []*account.AccountResponse{
		{
			AccountID:     "account-1",
			UserID:        testUserID,
			Name:          "Test Account 1",
			Type:          "savings",
			Currency:      "USD",
			AccountNumber: "1234567890",
			Issuer:        "Test Bank",
		},
		{
			AccountID:     "account-2",
			UserID:        testUserID,
			Name:          "Test Account 2",
			Type:          "checking",
			Currency:      "USD",
			AccountNumber: "0987654321",
			Issuer:        "Test Bank 2",
		},
	}

	testCases := []struct {
		name          string
		queryParams   string
		setupContext  func(ctx *gin.Context)
		buildStubs    func(usecase *mocks.AccountUsecase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			queryParams: "",
			setupContext: func(ctx *gin.Context) {
				ctx.Set("userID", testUserID)
			},
			buildStubs: func(usecase *mocks.AccountUsecase) {
				usecase.EXPECT().GetAccountsByUserID(testUserID, 10, 0).Return(testAccounts, 2, nil)
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
			buildStubs: func(usecase *mocks.AccountUsecase) {
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
			buildStubs: func(usecase *mocks.AccountUsecase) {
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
			buildStubs: func(usecase *mocks.AccountUsecase) {
				usecase.EXPECT().GetAccountsByUserID(testUserID, 10, 0).Return(nil, 0, errors.New("database error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewAccountUsecase(t)
			tc.buildStubs(mockUsecase)

			handler := account.NewAccountHandler(mockUsecase)

			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			engine := gin.New()
			engine.Use(middleware.ErrorHandler())

			engine.GET("/me/accounts", func(ctx *gin.Context) {
				tc.setupContext(ctx)
				handler.GetMyAccounts(ctx)
			})

			url := "/me/accounts" + tc.queryParams
			req, _ := http.NewRequest(http.MethodGet, url, nil)

			engine.ServeHTTP(recorder, req)

			tc.checkResponse(t, recorder)
		})
	}
}
