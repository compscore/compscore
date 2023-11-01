package engine

import (
	"context"
	"net/http"
	"time"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
)

// Start starts the engine
//
// @Summary Start the engine
// @Description Start the engine
// @Tags engine
// @Accept json
// @Produce json
// @Security ServiceAuth
// @Success 200 {object} models.Status
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /engine/start [post]
func Start(ctx *gin.Context) {
	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusUnauthorized,
			models.Error{
				Error: "user is not authenticate",
			},
		)
		return
	}

	if entTeam.Role != team.RoleAdmin {
		ctx.JSON(
			http.StatusUnauthorized,
			models.Error{
				Error: "request requires admin privileges",
			},
		)
		return
	}

	engineContext, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(config.Engine.Timeout)*time.Second,
	)
	defer cancel()

	message, err := client.Start(engineContext)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		models.Status{
			Message: message,
		},
	)
}
