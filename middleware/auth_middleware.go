package middleware

import (
	"net/http"
	"os"
	"user_service/helper"
	"user_service/models/web"
	"user_service/models/web/token"

	"github.com/dgrijalva/jwt-go"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) unauthorized(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	webResponse := web.Response{
		Code:   http.StatusUnauthorized,
		Status: "Unauthorized",
	}

	helper.WriteRequestToBody(w, webResponse)
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && (r.RequestURI == "/api/v1/user" || r.RequestURI == "/api/v1/auth") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		tokenAuth := r.Header.Get("Authorization")
		if tokenAuth == "" {
			middleware.unauthorized(w, r)
			return
		}

		var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
		claims := &token.TokenClaims{}

		token, err := jwt.ParseWithClaims(tokenAuth, claims, func(t *jwt.Token) (interface{}, error) {
			// ini harus mengembalikan kunci yang sama dengan yang kita buat
			return jwtTokenSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				middleware.unauthorized(w, r)
				return
			}
		}

		if !token.Valid {
			middleware.unauthorized(w, r)
			return
		}

		middleware.Handler.ServeHTTP(w, r)
	}
}
