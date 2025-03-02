package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/pkg/logger"
)

func HTTPErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			e := ctx.Errors.Last().Err

			var err *domain.Error

			if errors.As(e, &err) {
				logger.Error(err.Message)
				ctx.JSON(http.StatusInternalServerError, domain.HTTPResponse[string]{
					Severity: domain.SEVERITY_ERROR,
					Message:  err.Message,
					Body: func() *string {
						body := err.Error()
						return &body
					}(),
				})
			}
		}
	}
}
