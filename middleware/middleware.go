package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Use(route *gin.Engine, log *zap.Logger) {
	logger = log
	catchNotFoundErr(route)
	catchServiceErr(route)
}
