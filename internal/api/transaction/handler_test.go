package transaction_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/transaction"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/transaction/mocks"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const testUserID = "test-user-123"

func TestTransactionHandler_GetMyTransactions(t *testing.T) {
	testTransactions := []*transaction.TransactionResponse{
		{
			TransactionID: "txn-1",
			UserID:        testUserID,
			Name:          "Test Transaction 1",
			Image:         "image1.jpg",
			IsBank:        true,
		},
		{
			TransactionID: "txn-2",
			UserID:        testUserID,
			Name:          "Test Transaction 2",
			Image:         "image2.jpg",
			IsBank:        false,
		},
	}

	testCases := []struct {
		name          string
		queryParams   string
		setupContext  func(ctx *gin.Context)
		buildStubs    func(usecase *mocks.TransactionUsecase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			queryParams: "",
			setupContext: func(ctx *gin.Context) {
				ctx.Set("userID", testUserID)
			},
			buildStubs: func(usecase *mocks.TransactionUsecase) {
				usecase.EXPECT().GetTransactionsByUserID(testUserID, 10, 0).Return(testTransactions, 2, nil)
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
			buildStubs: func(usecase *mocks.TransactionUsecase) {
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
			buildStubs: func(usecase *mocks.TransactionUsecase) {
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
			buildStubs: func(usecase *mocks.TransactionUsecase) {
				usecase.EXPECT().GetTransactionsByUserID(testUserID, 10, 0).Return(nil, 0, errors.New("database error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewTransactionUsecase(t)
			tc.buildStubs(mockUsecase)

			handler := transaction.NewTransactionHandler(mockUsecase)

			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			engine := gin.New()
			engine.Use(middleware.ErrorHandler())

			engine.GET("/me/transactions", func(ctx *gin.Context) {
				tc.setupContext(ctx)
				handler.GetMyTransactions(ctx)
			})

			url := "/me/transactions" + tc.queryParams
			req, _ := http.NewRequest(http.MethodGet, url, nil)

			engine.ServeHTTP(recorder, req)

			tc.checkResponse(t, recorder)
		})
	}
}
