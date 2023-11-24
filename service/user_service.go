package service

import (
	"context"
	"user_service/models/web"
	"user_service/models/web/token"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UsersResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UsersResponse
	Delete(ctx context.Context, requestID string)
	FindAll(ctx context.Context) []web.UsersResponse
	FindById(ctx context.Context, userId string) web.UsersResponse
	Auth(ctx context.Context, request web.UserAuthRequest) token.TokenResponse
	WithRefreshToken(ctx context.Context, refreshToken string) token.TokenResponse
}
