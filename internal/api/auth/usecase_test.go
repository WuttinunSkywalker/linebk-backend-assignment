package auth_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/auth"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/user"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/user/mocks"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/config"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthUsecase_AuthUsecaseLoginWithPin(t *testing.T) {
	testUserID := "test-user-123"
	testPin := "123456"
	hashedPin, _ := bcrypt.GenerateFromPassword([]byte(testPin), bcrypt.DefaultCost)

	testUser := &user.User{
		UserID:  testUserID,
		PinHash: string(hashedPin),
	}

	testConfig := config.JWTConfig{
		Secret:               "test-secret",
		Issuer:               "test-issuer",
		AccessExpirySeconds:  3600,
		RefreshExpirySeconds: 86400,
	}

	testCases := []struct {
		name          string
		userID        string
		pin           string
		buildStubs    func(repo *mocks.UserRepository)
		checkResponse func(t *testing.T, response *auth.LoginResponse, err error)
	}{
		{
			name:   "OK",
			userID: testUserID,
			pin:    testPin,
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserByID(testUserID).Return(testUser, nil)
			},
			checkResponse: func(t *testing.T, response *auth.LoginResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.NotEmpty(t, response.AccessToken)
				require.NotEmpty(t, response.RefreshToken)
			},
		},
		{
			name:   "UserNotFound",
			userID: "nonexistent-user",
			pin:    testPin,
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserByID("nonexistent-user").Return(nil, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, response *auth.LoginResponse, err error) {
				require.Error(t, err)
				require.Nil(t, response)
				require.IsType(t, &errs.APIError{}, err)
				apiErr := err.(*errs.APIError)
				require.Equal(t, 401, apiErr.Code)
			},
		},
		{
			name:   "GetUserError",
			userID: "nonexistent-user",
			pin:    testPin,
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserByID("nonexistent-user").Return(nil, errors.New("some error"))
			},
			checkResponse: func(t *testing.T, response *auth.LoginResponse, err error) {
				require.Error(t, err)
				require.Nil(t, response)
				require.IsType(t, &errs.APIError{}, err)
				apiErr := err.(*errs.APIError)
				require.Equal(t, 500, apiErr.Code)
			},
		},
		{
			name:   "InvalidPin",
			userID: testUserID,
			pin:    "wrong-pin",
			buildStubs: func(repo *mocks.UserRepository) {
				repo.EXPECT().GetUserByID(testUserID).Return(testUser, nil)
			},
			checkResponse: func(t *testing.T, response *auth.LoginResponse, err error) {
				require.Error(t, err)
				require.Nil(t, response)
				require.IsType(t, &errs.APIError{}, err)
				apiErr := err.(*errs.APIError)
				require.Equal(t, 401, apiErr.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepository(t)
			tc.buildStubs(mockRepo)

			usecase := auth.NewAuthUsecase(mockRepo, testConfig)
			response, err := usecase.LoginWithPin(tc.userID, tc.pin)

			tc.checkResponse(t, response, err)
		})
	}
}
