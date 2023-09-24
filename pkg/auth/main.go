package auth

import (
	"fmt"
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Team string
	jwt.StandardClaims
}
type User struct {
	Username   string
	Expiration int
}

func GenerateJWT(team string) (string, int, error) {
	expiration := time.Now().Add(time.Duration(config.Web.Timeout) * time.Hour)

	claims := &Claims{
		Team: team,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(config.Web.JWTKey))

	return tokenStr, int(expiration.Unix()), err
}

func tokenParse(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(config.Web.JWTKey), nil
}

func ParseJWT(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, tokenParse)
	return token, claims, err
}

func Parse(ctx *gin.Context) (*ent.Team, error) {
	tokenString, err := ctx.Cookie("auth")
	if err != nil {
		return nil, fmt.Errorf("no auth cookie found")
	}

	_, claims, err := ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}

	team, err := data.Team.GetByName(claims.Team)
	if err != nil {
		return nil, err
	}

	return team, err
}
