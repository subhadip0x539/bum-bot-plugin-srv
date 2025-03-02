package services

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/ports"
)

type PluginServiceImpl struct {
	mongoRepo ports.MongoRepo
}

func (s *PluginServiceImpl) GetPlugins() ([]domain.Plugin, domain.Error) {
	var result []domain.Plugin

	if err := s.mongoRepo.FindAll("plugins", bson.M{}, &result); err != nil {
		return []domain.Plugin{{}}, domain.Error{
			Severity: domain.SEVERITY_ERROR,
			Message:  err.Error(),
			Error:    err,
		}
	}

	return result, domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Message:  "Success",
	}
}

func NewServerStatsService(mongoRepo ports.MongoRepo) *PluginServiceImpl {
	return &PluginServiceImpl{mongoRepo: mongoRepo}
}
