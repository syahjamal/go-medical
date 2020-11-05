package models

import "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}
