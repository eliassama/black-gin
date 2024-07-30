package middleware

import (
	"github.com/eliassama/black-gin/i18n"
	"github.com/eliassama/black-gin/res"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"runtime/debug"
)

func catchNotFoundErr(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		logger.Warn(
			"Route Not Found",
			zap.String("Method", ctx.Request.Method),
			zap.String("URL", ctx.Request.RequestURI),
			zap.Any("Header", ctx.Request.Header),
			zap.Any("Body", ctx.Request.Body),
			zap.Any("From", ctx.Request.Form),
			zap.String("Host", ctx.Request.Host),
		)

		res.GetMessage(i18n.SendNotFoundMsg(GetLanguage(ctx))).SendJson(ctx)
	})
}

func catchServiceErr(route *gin.Engine) {
	route.Use(gin.Recovery())

	route.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Sugar().Fatal(err)
				debug.PrintStack()
				res.GetMessage(i18n.SendInternalServerErrorMsg(GetLanguage(ctx))).SendJson(ctx)
				ctx.Abort()
			}
		}()

		ctx.Next()
	})
}
