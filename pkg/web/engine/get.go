package engine

import (
	"context"
	"net/http"
	"time"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// Get returns the status of the engine
//
// @Summary Status of the engine
// @Description Status of the engine
// @Tags engine
// @Accept json
// @Produce json
// @Success 200 {object} models.Status
// @Failure 500 {object} models.Error
// @Router /engine/status [get]
func Get(ctx *gin.Context) {
	if config.Production {
		cachedData, err := cache.Client.Get(ctx, "engine_status").Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		if err == nil {
			ctx.JSON(http.StatusOK, cachedData)
			return
		}
	}

	engineCtx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(config.Engine.Timeout)*time.Second,
	)
	defer cancel()

	status, message, err := client.Status(engineCtx)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
	}

	var statusString string

	switch status {
	case proto.StatusEnum_ERROR:
		statusString = "error"
	case proto.StatusEnum_RUNNING:
		statusString = "running"
	case proto.StatusEnum_PAUSED:
		statusString = "paused"
	case proto.StatusEnum_UNKNOWN:
		statusString = "unknown"
	}

	if config.Production {
		err = cache.Client.Set(ctx, "engine_status", statusString, config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}
	}

	ctx.JSON(
		http.StatusOK,
		models.EngineStatus{
			Status:  statusString,
			Message: message,
		},
	)

}
