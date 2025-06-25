package token_test

import (
	"testing"
	"time"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/token"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

const (
	testSecret = "test-secret-key"
	testIssuer = "test-issuer"
)

func TestGenerateAccessToken(t *testing.T) {
	userID := "test-user-123"

	testCases := []struct {
		name          string
		secretKey     []byte
		subject       string
		issuer        string
		expiresAt     time.Time
		checkResponse func(t *testing.T, token string, err error)
	}{
		{
			name:      "OK",
			secretKey: []byte(testSecret),
			subject:   userID,
			issuer:    testIssuer,
			expiresAt: time.Now().Add(time.Hour),
			checkResponse: func(t *testing.T, token string, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, token)
			},
		},
		{
			name:      "EmptySecret",
			secretKey: []byte(""),
			subject:   userID,
			issuer:    testIssuer,
			expiresAt: time.Now().Add(time.Hour),
			checkResponse: func(t *testing.T, token string, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, token)
			},
		},
		{
			name:      "EmptySubject",
			secretKey: []byte(testSecret),
			subject:   "",
			issuer:    testIssuer,
			expiresAt: time.Now().Add(time.Hour),
			checkResponse: func(t *testing.T, token string, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, token)
			},
		},
		{
			name:      "ExpiredToken",
			secretKey: []byte(testSecret),
			subject:   userID,
			issuer:    testIssuer,
			expiresAt: time.Now().Add(-time.Hour),
			checkResponse: func(t *testing.T, token string, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, token)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := token.GenerateAccessToken(tc.secretKey, tc.subject, tc.issuer, tc.expiresAt)
			tc.checkResponse(t, token, err)
		})
	}
}

func TestGenerateRefreshToken(t *testing.T) {
	userID := "test-user-123"

	testCases := []struct {
		name          string
		secretKey     []byte
		subject       string
		issuer        string
		expiresAt     time.Time
		checkResponse func(t *testing.T, token string, err error)
	}{
		{
			name:      "OK",
			secretKey: []byte(testSecret),
			subject:   userID,
			issuer:    testIssuer,
			expiresAt: time.Now().Add(24 * time.Hour),
			checkResponse: func(t *testing.T, token string, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, token)

				parsedToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(testSecret), nil
				})
				require.NoError(t, err)
				require.True(t, parsedToken.Valid)

				claims := parsedToken.Claims.(*jwt.RegisteredClaims)
				require.Equal(t, userID, claims.Subject)
				require.Equal(t, testIssuer, claims.Issuer)
			},
		},
		{
			name:      "EmptyIssuer",
			secretKey: []byte(testSecret),
			subject:   userID,
			issuer:    "",
			expiresAt: time.Now().Add(24 * time.Hour),
			checkResponse: func(t *testing.T, token string, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, token)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := token.GenerateRefreshToken(tc.secretKey, tc.subject, tc.issuer, tc.expiresAt)
			tc.checkResponse(t, token, err)
		})
	}
}
