package debit_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/debit"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/debit/mocks"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const testUserID = "test-user-123"

func TestDebitHandler_GetMyDebitCards(t *testing.T) {
	testDebitCards := []*debit.DebitCardResponse{
		{
			CardID: "card-1",
			UserID: testUserID,
			Name:   "Test Card 1",
		},
		{
			CardID: "card-2",
			UserID: testUserID,
			Name:   "Test Card 2",
		},
	}

	testCases := []struct {
		name          string
		queryParams   string
		setupContext  func(ctx *gin.Context)
		buildStubs    func(usecase *mocks.DebitUseCase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			queryParams: "",
			setupContext: func(ctx *gin.Context) {
				ctx.Set("userID", testUserID)
			},
			buildStubs: func(usecase *mocks.DebitUseCase) {
				usecase.EXPECT().GetDebitCardsByUserID(testUserID, 10, 0).Return(testDebitCards, 2, nil)
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
			buildStubs: func(usecase *mocks.DebitUseCase) {
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
			buildStubs: func(usecase *mocks.DebitUseCase) {
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
			buildStubs: func(usecase *mocks.DebitUseCase) {
				usecase.EXPECT().GetDebitCardsByUserID(testUserID, 10, 0).Return(nil, 0, errors.New("database error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := mocks.NewDebitUseCase(t)
			tc.buildStubs(mockUsecase)

			handler := debit.NewDebitHandler(mockUsecase)

			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			engine := gin.New()
			engine.Use(middleware.ErrorHandler())

			engine.GET("/me/debit-cards", func(ctx *gin.Context) {
				tc.setupContext(ctx)
				handler.GetMyDebitCards(ctx)
			})

			url := "/me/debit-cards" + tc.queryParams
			req, _ := http.NewRequest(http.MethodGet, url, nil)

			engine.ServeHTTP(recorder, req)

			tc.checkResponse(t, recorder)
		})
	}
}
