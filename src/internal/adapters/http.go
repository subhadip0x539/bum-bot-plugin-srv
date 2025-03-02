package adapters

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/handlers"
	"github.com/subhadip0x539/bum-bot-plugin-srv/src/internal/middlewares"
)

type HTTPClient struct {
	host   string
	port   int
	router *gin.Engine
}

func (c *HTTPClient) RegisterHandlers(pluginHandler handlers.PluginHandler) error {
	c.router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := c.router.Group("/v1")
	{
		plugins := v1.Group("/plugins")
		{
			plugins.GET("", pluginHandler.GetPlugins)
		}
	}

	return nil
}

func (c *HTTPClient) Run() error {
	return c.router.Run(fmt.Sprintf("%s:%d", c.host, c.port))
}

func NewHTTPClient(host string, port int, trustedProxies []string) (*HTTPClient, error) {
	router := gin.Default()

	router.Use(middlewares.HTTPErrorHandler())

	if err := router.SetTrustedProxies(trustedProxies); err != nil {
		return nil, err
	}

	return &HTTPClient{
		host:   host,
		port:   port,
		router: router,
	}, nil
}
