package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/core/ports"
)

type PluginHandler struct {
	svc ports.PluginService
}

func (h *PluginHandler) PatchStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *PluginHandler) GetPlugins(ctx *gin.Context) {
	plugins, err := h.svc.GetPlugins()
	if err.Details != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, domain.HTTPResponse[[]domain.Plugin]{
		Severity: domain.SEVERITY_SUCCESS,
		Body:     &plugins,
	})

}

func NewPluginHandler(svc ports.PluginService) *PluginHandler {
	return &PluginHandler{svc: svc}
}
