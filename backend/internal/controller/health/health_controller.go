package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusUnauthorized)
	})

	router.GET("/healthz", isHealthy)
}

func isHealthy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	c.Status(http.StatusOK)
}
