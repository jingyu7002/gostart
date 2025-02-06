package handlers

import (
	"net/http"
	"web-app/internal/config"
	"web-app/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, config.GetConfig().Version)
	})
	return r
}
