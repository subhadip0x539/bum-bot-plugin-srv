package services

import (
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
)

type PluginServiceImpl struct {
	mongoRepo ports.MongoRepo
}

func (s *PluginServiceImpl) GetPlugins() ([]domain.Plugin, *domain.Error) {
	var result []domain.Plugin

	if err := s.mongoRepo.FindAll("plugins", bson.M{}, &result); err != nil {
		return []domain.Plugin{{}}, domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}

	// return []domain.Plugin{{}}, domain.NewError(errors.New("error #01"), "Error #01", domain.SEVERITY_ERROR)

	return result, domain.NewError(nil, "", domain.SEVERITY_SUCCESS)
}

func NewServerStatsService(mongoRepo ports.MongoRepo) *PluginServiceImpl {
	return &PluginServiceImpl{mongoRepo: mongoRepo}
}
