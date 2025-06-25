package user

import (
	"net/http"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetMe(ctx *gin.Context)
	GetMyGreeting(ctx *gin.Context)
	GetUserPreview(ctx *gin.Context)
}

type userHandler struct {
	userUsecase UserUsecase
}

func NewUserHandler(userUsecase UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

// @Summary Get current user
// @Description Get user information of the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=user.UserResponse}
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Security BearerAuth
// @Router /api/users/me [get]
func (h *userHandler) GetMe(ctx *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		err := errs.Unauthorized("unauthorized")
		ctx.Error(err)
		return
	}

	user, err := h.userUsecase.GetUserByID(userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccess(user))
}

// @Summary Get user's greeting message
// @Description Get personalized greeting message for the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=user.UserGreetingResponse}
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Security BearerAuth
// @Router /api/users/me/greetings [get]
func (h *userHandler) GetMyGreeting(ctx *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		ctx.Error(errs.Unauthorized("unauthorized"))
		return
	}

	greeting, err := h.userUsecase.GetUserGreetingByUserID(userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccess(greeting))
}

// @Summary Get user's preview
// @Description Get preview information of a user
// @Tags users
// @Accept json
// @Produce json
// @Param userid path string true "User ID"
// @Success 200 {object} response.SuccessResponse{data=user.UserPreviewResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/users/{userid}/preview [get]
func (h *userHandler) GetUserPreview(ctx *gin.Context) {
	userID := ctx.Param("userid")
	if userID == "" {
		ctx.Error(errs.BadRequest("user ID is required"))
		return
	}

	userPreview, err := h.userUsecase.GetUserPreview(userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccess(userPreview))
}
