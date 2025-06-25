package user_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/user"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/user/mocks"
	"github.com/stretchr/testify/require"
)

func TestUserUsecase_GetUserByID(t *testing.T) {
	testUserID := "test-user-123"
	testTime := time.Now()
	testImage := "test-image.jpg"

	testUser := &user.User{
		UserID:       testUserID,
		Name:         "Test User",
		Image:        &testImage,
		PasswordHash: "hashed-password",
		PinHash:      "hashed-pin",
		CreatedAt:    testTime,
		UpdatedAt:    testTime,
	}

	testCases := []struct {
		name          string
		userID        string
		buildStubs    func(repo *mocks.UserRepository)
		checkResponse func(t *testing.T, response *user.UserResponse, err error)
	}{
		{
			name:   "OK",
			userID: testUserID,
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserByID(testUserID).Return(testUser, nil)
			},
			checkResponse: func(t *testing.T, response *user.UserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Equal(t, testUserID, response.UserID)
			},
		},
		{
			name:   "UserNotFound",
			userID: "non-existent-user",
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserByID("non-existent-user").Return(nil, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, response *user.UserResponse, err error) {
				require.Error(t, err)
				require.Nil(t, response)
			},
		},
		{
			name:   "DatabaseError",
			userID: testUserID,
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserByID(testUserID).Return(nil, errors.New("database connection error"))
			},
			checkResponse: func(t *testing.T, response *user.UserResponse, err error) {
				require.Error(t, err)
				require.Nil(t, response)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepository(t)
			tc.buildStubs(mockRepo)

			usecase := user.NewUserUsecase(mockRepo)
			response, err := usecase.GetUserByID(tc.userID)

			tc.checkResponse(t, response, err)
		})
	}
}

func TestUserUsecase_GetUserGreetingByUserID(t *testing.T) {
	testUserID := "test-user-123"
	testTime := time.Now()

	testGreeting := &user.UserGreeting{
		UserID:    testUserID,
		Greeting:  "Hello, World!",
		CreatedAt: testTime,
	}

	testCases := []struct {
		name          string
		userID        string
		buildStubs    func(repo *mocks.UserRepository)
		checkResponse func(t *testing.T, response *user.UserGreetingResponse, err error)
	}{
		{
			name:   "OK",
			userID: testUserID,
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserGreetingByUserID(testUserID).Return(testGreeting, nil)
			},
			checkResponse: func(t *testing.T, response *user.UserGreetingResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Equal(t, testUserID, response.UserID)
			},
		},
		{
			name:   "GreetingNotFound",
			userID: "non-existent-user",
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserGreetingByUserID("non-existent-user").Return(nil, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, response *user.UserGreetingResponse, err error) {
				require.Error(t, err)
				require.Nil(t, response)
			},
		},
		{
			name:   "DatabaseError",
			userID: testUserID,
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserGreetingByUserID(testUserID).Return(nil, errors.New("database connection error"))
			},
			checkResponse: func(t *testing.T, response *user.UserGreetingResponse, err error) {
				require.Error(t, err)
				require.Nil(t, response)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepository(t)
			tc.buildStubs(mockRepo)

			usecase := user.NewUserUsecase(mockRepo)
			response, err := usecase.GetUserGreetingByUserID(tc.userID)

			tc.checkResponse(t, response, err)
		})
	}
}

func TestUserUsecase_GetUserPreview(t *testing.T) {
	testUserID := "test-user-123"
	testTime := time.Now()
	testImage := "test-image.jpg"

	testUser := &user.User{
		UserID:       testUserID,
		Name:         "Test User",
		Image:        &testImage,
		PasswordHash: "hashed-password",
		PinHash:      "hashed-pin",
		CreatedAt:    testTime,
		UpdatedAt:    testTime,
	}

	testCases := []struct {
		name          string
		userID        string
		buildStubs    func(repo *mocks.UserRepository)
		checkResponse func(t *testing.T, response *user.UserPreviewResponse, err error)
	}{
		{
			name:   "OK",
			userID: testUserID,
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserByID(testUserID).Return(testUser, nil)
			},
			checkResponse: func(t *testing.T, response *user.UserPreviewResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Equal(t, testUser.Name, response.Name)
				require.Equal(t, testUser.Image, response.Image)
			},
		},
		{
			name:   "UserNotFound",
			userID: "non-existent-user",
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserByID("non-existent-user").Return(nil, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, response *user.UserPreviewResponse, err error) {
				require.Error(t, err)
				require.Nil(t, response)
			},
		},
		{
			name:   "DatabaseError",
			userID: testUserID,
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserByID(testUserID).Return(nil, errors.New("database connection error"))
			},
			checkResponse: func(t *testing.T, response *user.UserPreviewResponse, err error) {
				require.Error(t, err)
				require.Nil(t, response)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepository(t)
			tc.buildStubs(mockRepo)

			usecase := user.NewUserUsecase(mockRepo)
			response, err := usecase.GetUserPreview(tc.userID)

			tc.checkResponse(t, response, err)
		})
	}
}
