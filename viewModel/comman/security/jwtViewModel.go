package security

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	UserName string `json:"username"`
	ID string `json:"userId"`
	jwt.StandardClaims
}