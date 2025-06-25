package middleware

import (
	"strings"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/config"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware interface {
	ValidateToken() gin.HandlerFunc
}

type authMiddleware struct {
	cfg config.JWTConfig
}

func NewAuthMiddleware(cfg config.JWTConfig) AuthMiddleware {
	return &authMiddleware{
		cfg: cfg,
	}
}

func (m *authMiddleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(errs.Unauthorized("Authorization header is required"))
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.Error(errs.Unauthorized("Invalid token format"))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.Error(errs.Unauthorized("Token is required"))
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(m.cfg.Secret), nil
		})
		if err != nil {
			c.Error(errs.Unauthorized("Invalid token"))
			c.Abort()
			return
		}

		if !token.Valid {
			c.Error(errs.Unauthorized("Invalid token"))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			c.Error(errs.Unauthorized("Invalid claims"))
			c.Abort()
			return
		}

		if claims.Issuer != m.cfg.Issuer {
			c.Error(errs.Unauthorized("Invalid issuer"))
			c.Abort()
			return
		}

		c.Set("userID", claims.Subject)
		c.Next()
	}
}

func GetUserIDFromContext(c *gin.Context) (string, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return "", false
	}

	userIDStr, ok := userID.(string)
	return userIDStr, ok
}
