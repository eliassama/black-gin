package api

import (
	"github.com/eliassama/black-gin/api/route"
	"github.com/gin-gonic/gin"
)

type Enable = uint8

const (
	EnableHealthRoute Enable = iota + 1
	EnablePageRoute   Enable = iota + 1
)

func DefaultRelease(fn func(engine *gin.Engine), enableSlice []Enable) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engineCpp(engine, fn, enableSlice)
	return engine
}

func Default(fn func(engine *gin.Engine), enableSlice []Enable) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engineCpp(engine, fn, enableSlice)
	return engine
}

func engineCpp(engine *gin.Engine, fn func(engine *gin.Engine), enableSlice []Enable) {
	if len(enableSlice) != 0 {
		enableUsed := make(map[Enable]bool)

		for _, enable := range enableSlice {
			if _, exist := enableUsed[enable]; exist {
				continue
			}
			enableUsed[enable] = true

			switch enable {
			case EnableHealthRoute:
				route.LoadHealth(engine)
			case EnablePageRoute:
				route.LoadPage(engine)
			default:
			}

		}
	}

	fn(engine)

}
