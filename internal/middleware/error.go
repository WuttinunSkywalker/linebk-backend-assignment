package middleware

import (
	"errors"
	"net/http"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/logger"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/response"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors[0].Err

			var apiErr *errs.APIError
			if errors.As(err, &apiErr) {
				logger.Error(apiErr.Err)
				ctx.AbortWithStatusJSON(apiErr.Code, response.NewError(apiErr.Message))
				return
			} else {
				logger.Error(err)
				ctx.AbortWithStatusJSON(
					http.StatusInternalServerError,
					response.NewError("An unexpected internal server error occurred"),
				)
			}
		}

	}
}
