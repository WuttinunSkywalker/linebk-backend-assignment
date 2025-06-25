package debit_test

import (
	"errors"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/debit"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/debit/mocks"
	"github.com/stretchr/testify/require"
)

func TestDebitUsecase_GetDebitCardsByUserID(t *testing.T) {
	testUserID := "test-user-123"

	testDebitCards := []*debit.DebitCard{
		{
			CardID: "card-1",
			UserID: testUserID,
			Name:   "Test Card 1",
			DebitCardStatus: debit.DebitCardStatus{
				Status: "active",
			},
			DebitCardDetail: debit.DebitCardDetail{
				Issuer: "Test Bank",
				Number: "1234567890123456",
			},
			DebitCardDesign: debit.DebitCardDesign{
				Color:       "blue",
				BorderColor: "white",
			},
		},
		{
			CardID: "card-2",
			UserID: testUserID,
			Name:   "Test Card 2",
			DebitCardStatus: debit.DebitCardStatus{
				Status: "inactive",
			},
			DebitCardDetail: debit.DebitCardDetail{
				Issuer: "Test Bank 2",
				Number: "9876543210987654",
			},
			DebitCardDesign: debit.DebitCardDesign{
				Color:       "red",
				BorderColor: "black",
			},
		},
	}

	testCases := []struct {
		name          string
		userID        string
		limit         int
		offset        int
		buildStubs    func(repo *mocks.DebitRepository)
		checkResponse func(t *testing.T, response []*debit.DebitCardResponse, total int, err error)
	}{
		{
			name:   "OK",
			userID: testUserID,
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.DebitRepository) {
				repo.EXPECT().GetDebitCardsByUserID(testUserID, 10, 0).Return(testDebitCards, nil)
				repo.EXPECT().CountDebitCardsByUserID(testUserID).Return(2, nil)
			},
			checkResponse: func(t *testing.T, response []*debit.DebitCardResponse, total int, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Len(t, response, 2)
				require.Equal(t, 2, total)
			},
		},
		{
			name:   "GetCardsError",
			userID: testUserID,
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.DebitRepository) {
				repo.EXPECT().GetDebitCardsByUserID(testUserID, 10, 0).Return(nil, errors.New("database error"))
				repo.EXPECT().CountDebitCardsByUserID(testUserID).Return(0, nil)
			},
			checkResponse: func(t *testing.T, response []*debit.DebitCardResponse, total int, err error) {
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
			buildStubs: func(repo *mocks.DebitRepository) {
				repo.EXPECT().GetDebitCardsByUserID(testUserID, 10, 0).Return(testDebitCards, nil)
				repo.EXPECT().CountDebitCardsByUserID(testUserID).Return(0, errors.New("count error"))
			},
			checkResponse: func(t *testing.T, response []*debit.DebitCardResponse, total int, err error) {
				require.Error(t, err)
				require.Nil(t, response)
				require.Equal(t, 0, total)
			},
		},
		{
			name:   "EmptyResult",
			userID: "no-cards-user",
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.DebitRepository) {
				repo.EXPECT().GetDebitCardsByUserID("no-cards-user", 10, 0).Return([]*debit.DebitCard{}, nil)
				repo.EXPECT().CountDebitCardsByUserID("no-cards-user").Return(0, nil)
			},
			checkResponse: func(t *testing.T, response []*debit.DebitCardResponse, total int, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Len(t, response, 0)
				require.Equal(t, 0, total)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewDebitRepository(t)
			tc.buildStubs(mockRepo)

			usecase := debit.NewDebitUsecase(mockRepo)
			response, total, err := usecase.GetDebitCardsByUserID(tc.userID, tc.limit, tc.offset)

			tc.checkResponse(t, response, total, err)
		})
	}
}
