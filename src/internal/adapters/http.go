package adapters

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPClient struct {
	host   string
	port   int
	router *gin.Engine
}

func (c *HTTPClient) RegisterHandlers() error {
	c.router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return nil
}

func (c *HTTPClient) Run() error {
	return c.router.Run(fmt.Sprintf("%s:%d", c.host, c.port))
}

func NewHTTPClient(host string, port int, trustedProxies []string) (*HTTPClient, error) {
	router := gin.Default()

	if err := router.SetTrustedProxies(trustedProxies); err != nil {

		return nil, err
	}

	return &HTTPClient{
		host:   host,
		port:   port,
		router: router,
	}, nil
}
