package models

import "github.com/dgrijalva/jwt-go"

type TokenUser struct {
	ID       int64
	Username string
	jwt.StandardClaims
}
