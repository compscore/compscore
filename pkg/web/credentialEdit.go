package web

import (
	"io"
	"net/http"

	"encoding/json"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

type credential_edit_s struct {
	Password string `json:"password"`
}

func credential_edit(ctx *gin.Context) {
	var body credential_edit_s

	check := ctx.Param("check")

	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	if body.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No password provided",
		})
		return
	}

	_, err = data.Credential.UpdatePassword(entTeam.Number, check, body.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
