package route

import (
	"github.com/eliassama/black-gin/i18n"
	"github.com/eliassama/black-gin/middleware"
	"github.com/eliassama/black-gin/res"
	"github.com/gin-gonic/gin"
	"time"
)

// http 心跳检查
func ping(ctx *gin.Context) {
	middleware.SafeFunc(func() {
		res.GetMessage(i18n.SendOKMsg(middleware.GetLanguage(ctx))).SendJsonData(ctx, time.Now().Format("2006-01-02 15:04:05"))
	})
}
