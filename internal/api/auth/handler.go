package auth

import (
	"net/http"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	LoginWithPin(c *gin.Context)
}

type authHandler struct {
	authUsecase AuthUsecase
}

func NewAuthHandler(authUsecase AuthUsecase) AuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
	}
}

// @Summary Login with PIN
// @Description Authenticate user with user ID and PIN
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginWithPinRequest true "Login credentials"
// @Success 200 {object} response.SuccessResponse{data=LoginResponse} "Login successful"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/auth/login [post]
func (h *authHandler) LoginWithPin(c *gin.Context) {
	var req LoginWithPinRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errs.BadRequest(err.Error()))
		return
	}

	loginRes, err := h.authUsecase.LoginWithPin(req.UserID, req.Pin)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.NewSuccess(loginRes))
}
