package transaction

import (
	"net/http"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/pagination"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/response"
	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	GetMyTransactions(ctx *gin.Context)
}

type transactionHandler struct {
	transactionUsecase TransactionUsecase
}

func NewTransactionHandler(transactionUsecase TransactionUsecase) TransactionHandler {
	return &transactionHandler{
		transactionUsecase: transactionUsecase,
	}
}

// @Summary Get user's transactions
// @Description Get paginated list of transactions for the authenticated user
// @Tags transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} response.PaginatedResponse{data=[]transaction.TransactionResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Security BearerAuth
// @Router /api/transactions [get]
func (h *transactionHandler) GetMyTransactions(ctx *gin.Context) {
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
	transactions, total, err := h.transactionUsecase.GetTransactionsByUserID(userID, paginationParams.Limit, paginationParams.Offset())
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewPaginated(transactions, total, paginationParams.Page, paginationParams.Limit))
}
