package account

import (
	"net/http"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/pagination"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/response"
	"github.com/gin-gonic/gin"
)

type AccountHandler interface {
	GetMyAccounts(ctx *gin.Context)
}

type accountHandler struct {
	accountUsecase AccountUsecase
}

func NewAccountHandler(accountUsecase AccountUsecase) AccountHandler {
	return &accountHandler{
		accountUsecase: accountUsecase,
	}
}

// @Summary Get user's accounts
// @Description Get paginated list of accounts for the authenticated user
// @Tags accounts
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.PaginatedResponse{data=[]AccountResponse} "Successfully retrieved accounts"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /api/accounts [get]
func (h *accountHandler) GetMyAccounts(ctx *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		ctx.Error(errs.Unauthorized("unauthorized"))
		return
	}

	var paginationParams pagination.Params
	if err := ctx.ShouldBindQuery(&paginationParams); err != nil {
		ctx.Error(errs.BadRequest(err.Error()))
		return
	}

	paginationParams.Defaults()
	accounts, total, err := h.accountUsecase.GetAccountsByUserID(userID, paginationParams.Limit, paginationParams.Offset())
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewPaginated(accounts, total, paginationParams.Page, paginationParams.Limit))
}
