package route

import (
	"github.com/gin-gonic/gin"
)

func LoadHealth(engine *gin.Engine) {
	defaultGroup := engine.Group("/black.gin")
	healthGroup := defaultGroup.Group("/health")
	healthGroup.GET("/ping", ping)
	healthGroup.POST("/ping", ping)
	healthGroup.PUT("/ping", ping)
	healthGroup.DELETE("/ping", ping)
}

func LoadPage(engine *gin.Engine) {
	defaultGroup := engine.Group("/black.gin")
	pageGroup := defaultGroup.Group("/page")
	pageGroup.GET("/welcome", welcome)
	pageGroup.GET("/icon", icon)
}
