package app

import (
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/adapters"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/config"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/repositories"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/services"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/handlers"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/pkg/logger"
)

func Run(cfg config.Config) {
	http, err := adapters.NewHTTPClient(cfg.Server.Host, cfg.Server.Port, cfg.Server.TrustedProxies)
	if err != nil {
		logger.Fatal(err.Error())
	}

	mongo, err := adapters.NewMongoClient(cfg.Mongo.URI)
	if err != nil {
		logger.Error(err.Error())
	}

	if err := mongo.Connect(); err != nil {
		logger.Error(err.Error())
	}
	defer mongo.Disconnect()

	mongoRepo := repositories.NewMongoRepo(mongo.Client, cfg.Database)

	pluginService := services.NewServerStatsService(mongoRepo)

	pluginHandler := handlers.NewPluginHandler(pluginService)

	if err := http.RegisterHandlers(*pluginHandler); err != nil {
		logger.Fatal(err.Error())
	}

	http.Run()
}
