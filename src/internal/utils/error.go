package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/pkg/logger"
)

func LogEvent(err domain.Error, ctx *gin.Context) {
	switch err.Severity {
	case domain.SEVERITY_ERROR:
		logger.Error(err.Message)
		fmt.Print(err.Error)
		ctx.JSON(http.StatusInternalServerError, domain.HTTPResponse[interface{}]{
			Severity: domain.SEVERITY_ERROR,
			Message:  err.Message,
		})

	case domain.SEVERITY_WARNING:
		logger.Warn(err.Message)
	case domain.SEVERITY_SUCCESS:
		logger.Info(err.Message)
	default:
		logger.Info(err.Message)
	}
}
