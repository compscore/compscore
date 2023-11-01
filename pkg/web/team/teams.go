package team

import (
	"encoding/json"
	"net/http"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// Teams returns all teams
//
// @Summary Get all teams
// @Description Get all teams
// @Tags team
// @Accept json
// @Produce json
// @Success 200 {array} models.Team
// @Failure 500 {object} models.Error
// @Router /teams [get]
func Teams(ctx *gin.Context) {
	if config.Production {
		cachedData, err := cache.Client.Get(ctx, "teams").Result()
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
			ctx.String(http.StatusOK, cachedData)
			return
		}
	}

	entTeams, err := data.Team.GetAll()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	if config.Production {
		redisObject, err := json.Marshal(entTeams)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}
		err = cache.Client.Set(ctx, "teams", string(redisObject), config.Redis.SlowRefresh).Err()
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

	ctx.JSON(http.StatusOK, entTeams)
}
