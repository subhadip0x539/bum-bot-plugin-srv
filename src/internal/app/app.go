package app

import (
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/adapters"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/config"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/pkg/logger"
)

func Run(cfg config.Config) {
	http, err := adapters.NewHTTPClient(cfg.Server.Host, cfg.Server.Port, cfg.Server.TrustedProxies)

	if err != nil {
		logger.Fatal(err.Error())
	}

	if err := http.RegisterHandlers(); err != nil {
		logger.Fatal(err.Error())
	}

	http.Run()
}
