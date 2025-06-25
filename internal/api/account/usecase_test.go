package account_test

import (
	"errors"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/account"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/account/mocks"
	"github.com/stretchr/testify/require"
)

func TestAccountUsecase_GetAccountsByUserID(t *testing.T) {
	testUserID := "test-user-123"

	testAccounts := []*account.Account{
		{
			AccountID:     "account-1",
			UserID:        testUserID,
			Name:          "Test Account 1",
			Type:          "savings",
			Currency:      "USD",
			AccountNumber: "1234567890",
			Issuer:        "Test Bank",
			AccountBalance: account.AccountBalance{
				AccountID: "account-1",
				Amount:    1000.50,
			},
			AccountDetail: account.AccountDetail{
				AccountID:     "account-1",
				Color:         "blue",
				IsMainAccount: true,
				Progress:      85,
			},
		},
		{
			AccountID:     "account-2",
			UserID:        testUserID,
			Name:          "Test Account 2",
			Type:          "checking",
			Currency:      "USD",
			AccountNumber: "0987654321",
			Issuer:        "Test Bank 2",
			AccountBalance: account.AccountBalance{
				AccountID: "account-2",
				Amount:    500.25,
			},
			AccountDetail: account.AccountDetail{
				AccountID:     "account-2",
				Color:         "red",
				IsMainAccount: false,
				Progress:      60,
			},
		},
	}

	testCases := []struct {
		name          string
		userID        string
		limit         int
		offset        int
		buildStubs    func(repo *mocks.AccountRepository)
		checkResponse func(t *testing.T, response []*account.AccountResponse, total int, err error)
	}{
		{
			name:   "OK",
			userID: testUserID,
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.AccountRepository) {
				repo.EXPECT().GetAccountsByUserID(testUserID, 10, 0).Return(testAccounts, nil)
				repo.EXPECT().CountAccountsByUserID(testUserID).Return(2, nil)
			},
			checkResponse: func(t *testing.T, response []*account.AccountResponse, total int, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Len(t, response, 2)
				require.Equal(t, 2, total)
			},
		},
		{
			name:   "GetAccountsError",
			userID: testUserID,
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.AccountRepository) {
				repo.EXPECT().GetAccountsByUserID(testUserID, 10, 0).Return(nil, errors.New("database error"))
				repo.EXPECT().CountAccountsByUserID(testUserID).Return(0, nil)
			},
			checkResponse: func(t *testing.T, response []*account.AccountResponse, total int, err error) {
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
			buildStubs: func(repo *mocks.AccountRepository) {
				repo.EXPECT().GetAccountsByUserID(testUserID, 10, 0).Return(testAccounts, nil)
				repo.EXPECT().CountAccountsByUserID(testUserID).Return(0, errors.New("count error"))
			},
			checkResponse: func(t *testing.T, response []*account.AccountResponse, total int, err error) {
				require.Error(t, err)
				require.Nil(t, response)
				require.Equal(t, 0, total)
			},
		},
		{
			name:   "EmptyResult",
			userID: "no-accounts-user",
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.AccountRepository) {
				repo.EXPECT().GetAccountsByUserID("no-accounts-user", 10, 0).Return([]*account.Account{}, nil)
				repo.EXPECT().CountAccountsByUserID("no-accounts-user").Return(0, nil)
			},
			checkResponse: func(t *testing.T, response []*account.AccountResponse, total int, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Len(t, response, 0)
				require.Equal(t, 0, total)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewAccountRepository(t)
			tc.buildStubs(mockRepo)

			usecase := account.NewAccountUsecase(mockRepo)
			response, total, err := usecase.GetAccountsByUserID(tc.userID, tc.limit, tc.offset)

			tc.checkResponse(t, response, total, err)
		})
	}
}
