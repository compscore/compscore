package web

import (
	"fmt"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/compscore/compscore/pkg/web/admin"
	"github.com/compscore/compscore/pkg/web/check"
	"github.com/compscore/compscore/pkg/web/credential"
	"github.com/compscore/compscore/pkg/web/engine"
	"github.com/compscore/compscore/pkg/web/round"
	"github.com/compscore/compscore/pkg/web/scoreboard"
	"github.com/compscore/compscore/pkg/web/status"
	"github.com/compscore/compscore/pkg/web/team"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
	API    *gin.RouterGroup
)

func Start() {
	if config.Web.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.Default()
	Router.SetTrustedProxies(nil)

	cors_urls := []string{
		fmt.Sprintf("http://%s:%d", config.Web.Hostname, config.Web.Port),
		fmt.Sprintf("https://%s:%d", config.Web.Hostname, config.Web.Port),
		fmt.Sprintf("http://%s:3000", config.Web.Hostname),
		fmt.Sprintf("https://%s:3000", config.Web.Hostname),
		fmt.Sprintf("http://%s:5173", config.Web.Hostname),
		fmt.Sprintf("https://%s:5173", config.Web.Hostname),
	}

	Router.Use(cors.New(cors.Config{
		AllowOrigins:     cors_urls,
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	Router.Use(auth.JWTMiddleware)

	API = Router.Group("/api")

	LoadRoutes()

	client.Open()
	defer client.Close()

	Router.Run(fmt.Sprintf(":%d", config.Web.Port))
}

func LoadRoutes() {
	// General Endpoints
	API.POST("/login", login)
	API.POST("/info", info)
	API.POST("/password", password)

	// Admin Endpoints
	API.POST("/admin/password", admin.Password)
	API.POST("/admin/login", admin.Login)

	// Engine Endpoints
	API.GET("/engine", engine.Get)
	API.POST("/engine/start", engine.Start)
	API.POST("/engine/stop", engine.Stop)

	// Scoreboard Endpoints
	API.GET("/scoreboard", scoreboard.Scoreboard)
	API.GET("/scoreboard/round/:round", scoreboard.Round)
	API.GET("/scoreboard/team/:team", scoreboard.Team)
	API.GET("/scoreboard/team/:team/:round", scoreboard.TeamRound)
	API.GET("/scoreboard/check/:check", scoreboard.Check)
	API.GET("/scoreboard/check/:check/:round", scoreboard.CheckRound)
	API.GET("/scoreboard/status/:team/:check", scoreboard.Status)
	API.GET("/scoreboard/status/:team/:check/:round", scoreboard.StatusRound)

	// Credentials Endpoints
	API.GET("/credentials", credential.Credentials)
	API.GET("/credential/:check", credential.Get)
	API.POST("/credential/:check", credential.Post)

	// Check Endpoints
	API.GET("/checks", check.Checks)
	API.GET("/check/:check", check.Get)

	// Round Endpoints
	API.GET("/rounds", round.Rounds)
	API.GET("/round/latest", round.Latest)
	API.GET("/round/:round", round.Get)

	// Status Endpoints
	API.GET("/statuses", status.Statuses)
	API.GET("/status/:team/:check/:round", status.Get)
	API.GET("/status/team/:team", status.GetByTeam)
	API.GET("/status/check/:check", status.GetByCheck)
	API.GET("/status/round/:round", status.GetByRound)
	API.GET("/status/team/:team/check/:check", status.GetByTeamCheck)
	API.GET("/status/team/:team/round/:round", status.GetByTeamRound)
	API.GET("/status/check/:check/round/:round", status.GetByCheckRound)

	// Team Endpoints
	API.GET("/teams", team.Teams)
	API.GET("/team/:team", team.Get)
}
