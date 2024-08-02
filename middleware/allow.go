package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DefaultCors  默认跨域处理
//
//goland:noinspection ALL
func DefaultCors(origin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if origin == "" {
			origin = c.Request.Header.Get("Origin")
		}

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "false")
		c.Set("content-type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

// AllowOrigin 仅做跨域处理
//
//goland:noinspection ALL
func AllowOrigin(origin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if origin == "" {
			origin = c.Request.Header.Get("Origin")
		}

		c.Header("Access-Control-Allow-Origin", origin)

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

// AllowHeaders 允许的请求头
//
//goland:noinspection ALL
func AllowHeaders(headers map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for key, value := range headers {
			c.Header(key, value)
		}
		c.Next()
	}
}
