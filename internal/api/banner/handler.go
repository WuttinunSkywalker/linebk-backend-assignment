package banner

import (
	"net/http"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/pagination"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/response"
	"github.com/gin-gonic/gin"
)

type BannerHandler interface {
	GetMyBanners(ctx *gin.Context)
}

type bannerHandler struct {
	bannerUsecase BannerUsecase
}

func NewBannerHandler(bannerUsecase BannerUsecase) BannerHandler {
	return &bannerHandler{
		bannerUsecase: bannerUsecase,
	}
}

// @Summary Get user's banners
// @Description Get paginated list of banners for the authenticated user
// @Tags banners
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} response.PaginatedResponse{data=[]banner.BannerResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Security BearerAuth
// @Router /api/banners [get]
func (h *bannerHandler) GetMyBanners(ctx *gin.Context) {
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
	banners, total, err := h.bannerUsecase.GetBannersByUserID(userID, paginationParams.Limit, paginationParams.Offset())
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewPaginated(banners, total, paginationParams.Page, paginationParams.Limit))
}
