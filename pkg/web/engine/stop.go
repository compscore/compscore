package engine

import (
	"context"
	"net/http"
	"time"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/gin-gonic/gin"
)

func Stop(ctx *gin.Context) {
	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if entTeam.Role != team.RoleAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	engineContext, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(config.Engine.Timeout)*time.Second,
	)
	defer cancel()

	message, err := client.Pause(engineContext)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
