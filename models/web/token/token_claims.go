package token

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UserId    string
	Email     string
	FirstName string
	LastName  string
	jwt.StandardClaims
}
