package middleware

import (
	"github.com/gin-gonic/gin"
)

// GetLanguage 从 gin http 上下文中获取语言
func GetLanguage(ctx *gin.Context) string {
	languageName := ctx.GetHeader("Accept-Language")
	return languageName
}
