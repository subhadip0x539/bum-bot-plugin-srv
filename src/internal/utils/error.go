package utils

import (
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/pkg/logger"
)

func LogEvent(err domain.Error) {
	switch err.Severity {
	case domain.SEVERITY_ERROR:
		logger.Error(err.Message)
	case domain.SEVERITY_WARNING:
		logger.Warn(err.Message)
	case domain.SEVERITY_SUCCESS:
		logger.Info(err.Message)
	default:
		logger.Info(err.Message)
	}
}
