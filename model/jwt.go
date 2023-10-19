package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomJwtClaims struct {
	Name string
	ID UserID
	jwt.RegisteredClaims
}

type CustomRefreshJwtClaims struct {
	ID UserID
	jwt.RegisteredClaims
}