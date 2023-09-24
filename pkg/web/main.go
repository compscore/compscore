package web

import (
	"fmt"

	"github.com/compscore/compscore/pkg/config"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
	API    *gin.RouterGroup
)

func Start() {
	// gin.SetMode(gin.ReleaseMode)

	Router = gin.Default()

	Router.SetTrustedProxies(nil)
	fmt.Println(config.Web.APIPath)
	API = Router.Group(config.Web.APIPath)

	LoadRoutes()

	Router.Run(fmt.Sprintf("%s:%d", config.Web.Hostname, config.Web.Port))
}

func LoadRoutes() {
	API.POST("/login", login)
}
