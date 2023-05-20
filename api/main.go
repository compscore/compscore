package api

import (
	"github.com/compscore/compscore/config"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()

	Router.SetTrustedProxies(nil)
}

func Start() {
	loadRoutes()
	Router.Run(config.Hostname + ":" + config.Port)
}

func loadRoutes() {
}
