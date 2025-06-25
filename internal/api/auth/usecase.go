package auth

import (
	"database/sql"
	"errors"
	"time"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/user"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/config"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	LoginWithPin(userID, pin string) (*LoginResponse, error)
}

type authUsecase struct {
	user user.UserRepository
	cfg  config.JWTConfig
}

func NewAuthUsecase(user user.UserRepository, cfg config.JWTConfig) AuthUsecase {
	return &authUsecase{
		user: user,
		cfg:  cfg,
	}
}

func (u *authUsecase) LoginWithPin(userID, pin string) (*LoginResponse, error) {
	user, err := u.user.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.Unauthorized("invalid credentials")
		}

		return nil, errs.Internal(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PinHash), []byte(pin))
	if err != nil {
		return nil, errs.Unauthorized("invalid credentials")
	}

	expiresAt := time.Now().Add(time.Duration(u.cfg.AccessExpirySeconds) * time.Second)
	accessToken, err := token.GenerateAccessToken([]byte(u.cfg.Secret), userID, u.cfg.Issuer, expiresAt)
	if err != nil {
		return nil, errs.Internal(err)
	}

	expiresAt = time.Now().Add(time.Duration(u.cfg.RefreshExpirySeconds) * time.Second)
	refreshToken, err := token.GenerateRefreshToken([]byte(u.cfg.Secret), userID, u.cfg.Issuer, expiresAt)
	if err != nil {
		return nil, errs.Internal(err)
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
