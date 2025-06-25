package banner_test

import (
	"errors"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/banner"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/banner/mocks"
	"github.com/stretchr/testify/require"
)

func TestBannerUseCase_GetBannersByUserID(t *testing.T) {
	testUserID := "test-user-123"

	testBanners := []*banner.Banner{
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
		userID        string
		limit         int
		offset        int
		buildStubs    func(repo *mocks.BannerRepository)
		checkResponse func(t *testing.T, response []*banner.BannerResponse, total int, err error)
	}{
		{
			name:   "OK",
			userID: testUserID,
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.BannerRepository) {
				repo.EXPECT().GetBannersByUserID(testUserID, 10, 0).Return(testBanners, nil)
				repo.EXPECT().CountBannersByUserID(testUserID).Return(2, nil)
			},
			checkResponse: func(t *testing.T, response []*banner.BannerResponse, total int, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Len(t, response, 2)
				require.Equal(t, 2, total)
			},
		},
		{
			name:   "GetBannersError",
			userID: testUserID,
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.BannerRepository) {
				repo.EXPECT().GetBannersByUserID(testUserID, 10, 0).Return(nil, errors.New("database error"))
				repo.EXPECT().CountBannersByUserID(testUserID).Return(0, nil)
			},
			checkResponse: func(t *testing.T, response []*banner.BannerResponse, total int, err error) {
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
			buildStubs: func(repo *mocks.BannerRepository) {
				repo.EXPECT().GetBannersByUserID(testUserID, 10, 0).Return(testBanners, nil)
				repo.EXPECT().CountBannersByUserID(testUserID).Return(0, errors.New("count error"))
			},
			checkResponse: func(t *testing.T, response []*banner.BannerResponse, total int, err error) {
				require.Error(t, err)
				require.Nil(t, response)
				require.Equal(t, 0, total)
			},
		},
		{
			name:   "EmptyResult",
			userID: "no-banners-user",
			limit:  10,
			offset: 0,
			buildStubs: func(repo *mocks.BannerRepository) {
				repo.EXPECT().GetBannersByUserID("no-banners-user", 10, 0).Return([]*banner.Banner{}, nil)
				repo.EXPECT().CountBannersByUserID("no-banners-user").Return(0, nil)
			},
			checkResponse: func(t *testing.T, response []*banner.BannerResponse, total int, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Len(t, response, 0)
				require.Equal(t, 0, total)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewBannerRepository(t)
			tc.buildStubs(mockRepo)

			usecase := banner.NewBannerUsecase(mockRepo)
			response, total, err := usecase.GetBannersByUserID(tc.userID, tc.limit, tc.offset)

			tc.checkResponse(t, response, total, err)
		})
	}
}
