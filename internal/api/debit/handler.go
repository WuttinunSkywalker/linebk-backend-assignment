package debit

import (
	"net/http"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/pagination"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/response"
	"github.com/gin-gonic/gin"
)

type DebitHandler interface {
	GetMyDebitCards(ctx *gin.Context)
}

type debitHandler struct {
	debitUsecase DebitUseCase
}

func NewDebitHandler(debitUsecase DebitUseCase) DebitHandler {
	return &debitHandler{
		debitUsecase: debitUsecase,
	}
}

// @Summary Get user's debit cards
// @Description Get paginated list of debit cards for the authenticated user
// @Tags debit-cards
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} response.PaginatedResponse{data=[]debit.DebitCardResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Security BearerAuth
// @Router /api/debits [get]
func (h *debitHandler) GetMyDebitCards(ctx *gin.Context) {
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
	debitCards, total, err := h.debitUsecase.GetDebitCardsByUserID(userID, paginationParams.Limit, paginationParams.Offset())
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewPaginated(debitCards, total, paginationParams.Page, paginationParams.Limit))
}
