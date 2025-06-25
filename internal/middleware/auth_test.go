package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/config"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

const (
	authorizationTypeBearer = "Bearer"
	testSecret              = "test-secret-key-for-jwt-signing"
	testIssuer              = "test-issuer"
)

func TestValidateToken(t *testing.T) {
	userID := "test-user-123"

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request) {
				addAuthorization(t, request, authorizationTypeBearer, userID, time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidAuthorizationFormat",
			setupAuth: func(t *testing.T, request *http.Request) {
				request.Header.Set("Authorization", "InvalidFormat token-here")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "EmptyToken",
			setupAuth: func(t *testing.T, request *http.Request) {
				request.Header.Set("Authorization", "Bearer ")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "ExpiredToken",
			setupAuth: func(t *testing.T, request *http.Request) {
				addAuthorization(t, request, authorizationTypeBearer, userID, -time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidSigningMethod",
			setupAuth: func(t *testing.T, request *http.Request) {
				token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
					Subject:   userID,
					Issuer:    testIssuer,
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
					IssuedAt:  jwt.NewNumericDate(time.Now()),
				})

				request.Header.Set("Authorization", "Bearer "+token.Raw)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidIssuer",
			setupAuth: func(t *testing.T, request *http.Request) {
				addAuthorizationWithIssuer(t, request, authorizationTypeBearer, userID, "wrong-issuer", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "MalformedToken",
			setupAuth: func(t *testing.T, request *http.Request) {
				request.Header.Set("Authorization", "Bearer invalid.malformed.token")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			cfg := config.JWTConfig{
				Secret: testSecret,
				Issuer: testIssuer,
			}

			mockMiddleware := middleware.NewAuthMiddleware(cfg)
			errorMiddleware := middleware.ErrorHandler()

			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.Use(errorMiddleware)

			authPath := "/auth"
			router.GET(authPath, mockMiddleware.ValidateToken(), func(ctx *gin.Context) {
				if userID, exists := middleware.GetUserIDFromContext(ctx); exists {
					ctx.JSON(http.StatusOK, gin.H{"userID": userID})
				} else {
					ctx.JSON(http.StatusOK, gin.H{})
				}
			})

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request)
			router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetUserIDFromContext(t *testing.T) {
	testCases := []struct {
		name           string
		setupContext   func() *gin.Context
		expectedUserID string
		expectedExists bool
	}{
		{
			name: "UserIDExists",
			setupContext: func() *gin.Context {
				gin.SetMode(gin.TestMode)
				ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
				ctx.Set("userID", "test-user-123")
				return ctx
			},
			expectedUserID: "test-user-123",
			expectedExists: true,
		},
		{
			name: "UserIDNotExists",
			setupContext: func() *gin.Context {
				gin.SetMode(gin.TestMode)
				ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
				return ctx
			},
			expectedUserID: "",
			expectedExists: false,
		},
		{
			name: "UserIDWrongType",
			setupContext: func() *gin.Context {
				gin.SetMode(gin.TestMode)
				ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
				ctx.Set("userID", 123)
				return ctx
			},
			expectedUserID: "",
			expectedExists: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tc.setupContext()

			userID, exists := middleware.GetUserIDFromContext(ctx)

			require.Equal(t, tc.expectedUserID, userID)
			require.Equal(t, tc.expectedExists, exists)
		})
	}
}

func addAuthorization(t *testing.T, request *http.Request, authorizationType string, userID string, duration time.Duration) {
	addAuthorizationWithIssuer(t, request, authorizationType, userID, testIssuer, duration)
}

func addAuthorizationWithIssuer(t *testing.T, request *http.Request, authorizationType string, userID string, issuer string, duration time.Duration) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userID,
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	tokenString, err := token.SignedString([]byte(testSecret))
	require.NoError(t, err)

	authorizationHeader := authorizationType + " " + tokenString
	request.Header.Set("Authorization", authorizationHeader)
}
