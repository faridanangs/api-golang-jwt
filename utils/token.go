package utils

import (
	"os"
	"time"
	"user_service/exception"
	"user_service/helper"
	"user_service/models/web/token"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(request token.TokenCreateRequest, value time.Duration) string {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	var expiredToken = time.Now().Add(value * time.Minute)

	claims := &token.TokenClaims{
		UserId:    request.UserId,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredToken.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(jwtTokenSecret)
	helper.PanicError(err)

	return stringToken
}

func ClaimsToken(userToken string) *token.TokenClaims {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	claims := &token.TokenClaims{}

	token, err := jwt.ParseWithClaims(userToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtTokenSecret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			panic(exception.NewErrorUnauthorized(err.Error()))
		}
	}

	if !token.Valid {
		panic(exception.NewErrorUnauthorized(err.Error()))
	}

	return claims

}
