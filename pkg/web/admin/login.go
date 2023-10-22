package admin

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/gin-gonic/gin"
)

type login_s struct {
	Team string `json:"team"`
}

func Login(ctx *gin.Context) {
	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	if entTeam.Role != team.RoleAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	var body login_s

	body_bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = json.Unmarshal(body_bytes, &body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, expiration, err := auth.GenerateJWT(body.Team)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"name":       "auth",
			"token":      token,
			"expiration": expiration,
			"path":       "/",
			"domain":     config.Web.Hostname,
			"secure":     false,
			"httponly":   true,
		},
	)
}
