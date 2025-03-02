package ports

import "github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/domain"

type PluginService interface {
	GetPlugins() ([]domain.Plugin, domain.Error)
}
