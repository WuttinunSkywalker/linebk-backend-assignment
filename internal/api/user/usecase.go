package user

import (
	"database/sql"
	"errors"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
)

type UserUsecase interface {
	GetUserByID(id string) (*UserResponse, error)
	GetUserGreetingByUserID(id string) (*UserGreetingResponse, error)
	GetUserPreview(id string) (*UserPreviewResponse, error)
}

type userUsecase struct {
	user UserRepository
}

func NewUserUsecase(user UserRepository) UserUsecase {
	return &userUsecase{
		user: user,
	}
}

func (u *userUsecase) GetUserByID(id string) (*UserResponse, error) {
	user, err := u.user.GetUserByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound("user not found")
		}

		return nil, errs.Internal(err)
	}

	return NewUserResponse(user), nil
}

func (u *userUsecase) GetUserGreetingByUserID(id string) (*UserGreetingResponse, error) {
	greeting, err := u.user.GetUserGreetingByUserID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound("greeting not found for user")
		}

		return nil, errs.Internal(err)
	}

	return NewUserGreetingResponse(greeting), nil
}

func (u *userUsecase) GetUserPreview(id string) (*UserPreviewResponse, error) {
	user, err := u.user.GetUserByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound("user not found")
		}

		return nil, errs.Internal(err)
	}

	return NewUserPreviewResponse(user), nil
}
