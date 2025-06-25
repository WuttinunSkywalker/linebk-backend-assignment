package transaction_test

import (
	"errors"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/transaction"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/transaction/mocks"
	"github.com/stretchr/testify/require"
)

func TestTransactionUseCase_GetTransactionsByUserID(t *testing.T) {
	testUserID := "test-user-123"

	testTransactions := []*transaction.Transaction{
		{
			TransactionID: "txn-1",
			UserID:        testUserID,
			Name:          "Test Transaction 1",
			Image:         "image1.jpg",
			IsBank:        true,
			CreatedAt:     "2023-01-01T00:00:00Z",
		},
		{
			TransactionID: "txn-2",
			UserID:        testUserID,
			Name:          "Test Transaction 2",
			Image:         "image2.jpg",
			IsBank:        false,
			CreatedAt:     "2023-01-02T00:00:00Z",
		},
	}

	testCases := []struct {
		name          string
		userID        string
		limit         int
		offset        int
		buildStubs    func(repo *mocks.TransactionRepository)
		checkResponse func(t *testing.T, response []*transaction.TransactionResponse, total int, err error)
	}{
		{
			name:   "OK",
			userID: testUserID,
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.TransactionRepository) {
				repo.EXPECT().GetTransactionsByUserID(testUserID, 10, 0).Return(testTransactions, nil)
				repo.EXPECT().CountTransactionsByUserID(testUserID).Return(2, nil)
			},
			checkResponse: func(t *testing.T, response []*transaction.TransactionResponse, total int, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Len(t, response, 2)
				require.Equal(t, 2, total)
			},
		},
		{
			name:   "GetTransactionsError",
			userID: testUserID,
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.TransactionRepository) {
				repo.EXPECT().GetTransactionsByUserID(testUserID, 10, 0).Return(nil, errors.New("database error"))
				repo.EXPECT().CountTransactionsByUserID(testUserID).Return(0, nil)
			},
			checkResponse: func(t *testing.T, response []*transaction.TransactionResponse, total int, err error) {
				require.Error(t, err)
				require.Nil(t, response)
				require.Equal(t, 0, total)
			},
		},
		{
			name:   "CountError",
			userID: testUserID,
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.TransactionRepository) {
				repo.EXPECT().GetTransactionsByUserID(testUserID, 10, 0).Return(testTransactions, nil)
				repo.EXPECT().CountTransactionsByUserID(testUserID).Return(0, errors.New("count error"))
			},
			checkResponse: func(t *testing.T, response []*transaction.TransactionResponse, total int, err error) {
				require.Error(t, err)
				require.Nil(t, response)
				require.Equal(t, 0, total)
			},
		},
		{
			name:   "EmptyResult",
			userID: "no-transactions-user",
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.TransactionRepository) {
				repo.EXPECT().GetTransactionsByUserID("no-transactions-user", 10, 0).Return([]*transaction.Transaction{}, nil)
				repo.EXPECT().CountTransactionsByUserID("no-transactions-user").Return(0, nil)
			},
			checkResponse: func(t *testing.T, response []*transaction.TransactionResponse, total int, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Len(t, response, 0)
				require.Equal(t, 0, total)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewTransactionRepository(t)
			tc.buildStubs(mockRepo)

			usecase := transaction.NewTransactionUsecase(mockRepo)
			response, total, err := usecase.GetTransactionsByUserID(tc.userID, tc.limit, tc.offset)

			tc.checkResponse(t, response, total, err)
		})
	}
}
