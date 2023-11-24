package service

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"
	"user_service/exception"
	"user_service/helper"
	"user_service/models/entity"
	"user_service/models/web"
	"user_service/models/web/token"
	"user_service/repository"
	"user_service/utils"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceIplm struct {
	Validate       *validator.Validate
	DB             *sql.DB
	UserRepository repository.UserRepositry
}

func NewUserServiceIplm(userRepository repository.UserRepositry, validate *validator.Validate, db *sql.DB) UserService {
	return &UserServiceIplm{
		Validate:       validate,
		DB:             db,
		UserRepository: userRepository,
	}
}

func (service *UserServiceIplm) Create(ctx context.Context, request web.UserCreateRequest) web.UsersResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	// encript password use crypto
	passwordHash, err := utils.HassPasword(request.Password)
	helper.PanicError(err)

	user := entity.Users{
		Id:        utils.Uuid(),
		Password:  passwordHash,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return helper.UserResponse(user)

}
func (service *UserServiceIplm) Update(ctx context.Context, request web.UserUpdateRequest) web.UsersResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.UpdatedAt = time.Now().Unix()

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.UserResponse(user)

}
func (service *UserServiceIplm) Delete(ctx context.Context, requestID string) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, requestID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, user)
}
func (service *UserServiceIplm) FindById(ctx context.Context, userId string) web.UsersResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.UserResponse(user)
}
func (service *UserServiceIplm) FindAll(ctx context.Context) []web.UsersResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	return helper.UserResponses(users)
}

func (service *UserServiceIplm) Auth(ctx context.Context, request web.UserAuthRequest) token.TokenResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	helper.PanicError(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	
	if err != nil {
		panic(exception.NewErrorUnauthorized(err.Error()))
	}

	jwtExpiredTimeToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_TOKEN"))
	jwtExpiredTimeRefreshToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_REFRESH_TOKEN"))

	tokenCreateRequest := &token.TokenCreateRequest{
		UserId:    user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	tokenResponse := token.TokenResponse{
		Token:        utils.CreateToken(*tokenCreateRequest, time.Duration(jwtExpiredTimeToken)),
		TokenRefresh: utils.CreateToken(*tokenCreateRequest, time.Duration(jwtExpiredTimeRefreshToken)),
	}

	return tokenResponse

}

func (service *UserServiceIplm) WithRefreshToken(ctx context.Context, refreshToken string) token.TokenResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	claims := utils.ClaimsToken(refreshToken)

	_, err = service.UserRepository.FindById(ctx, tx, claims.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	tokenCreateRequest := token.TokenCreateRequest{
		UserId:    claims.UserId,
		Email:     claims.Email,
		FirstName: claims.FirstName,
		LastName:  claims.LastName,
	}

	jwtExpiredTimeToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_TOKEN"))
	jwtExpiredTimeRefreshToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_REFRESH_TOKEN"))

	tokenRefresh := token.TokenResponse{
		Token:        utils.CreateToken(tokenCreateRequest, time.Duration(jwtExpiredTimeToken)),
		TokenRefresh: utils.CreateToken(tokenCreateRequest, time.Duration(jwtExpiredTimeRefreshToken)),
	}

	return tokenRefresh
}
