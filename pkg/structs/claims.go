package structs

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Team string
	jwt.RegisteredClaims
}

type contextKey struct {
	name string
}

var TEAM_CTX_KEY = &contextKey{"team"}
