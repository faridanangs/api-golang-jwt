package main

import (
	"fmt"
	"net/http"
	"user_service/app"
	"user_service/controller"
	"user_service/exception"
	"user_service/helper"
	"user_service/middleware"
	"user_service/repository"
	"user_service/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// load environment
	envErr := godotenv.Load(".env")
	helper.PanicError(envErr)

	// Initialize
	router := httprouter.New()

	// connect to database
	db := app.Database()

	// validate
	validate := validator.New()

	// Users
	userRepository := repository.NewRepositoryIplm()
	userService := service.NewUserServiceIplm(userRepository, validate, db)
	userController := controller.NewUserControllerIplm(userService)

	// User Router
	router.POST("/api/v1/user", userController.Create)
	router.POST("/api/v1/auth", userController.Auth)
	router.POST("/api/v1/refresh-token", userController.WithRefreshToken)
	router.PUT("/api/v1/user/:user_id", userController.Update)
	router.DELETE("/api/v1/user/:user_id", userController.Delete)
	router.GET("/api/v1/user/:user_id", userController.FindById)
	router.GET("/api/v1/user", userController.FindAll)

	// Images
	imageRepository := repository.NewImagesRepositoryIplm()
	imagesService := service.NewImagesServiceIplm(db, *validate, imageRepository)
	imagesController := controller.NewImagesControllerIplm(imagesService)

	// images Router
	router.POST("/api/v1/images", imagesController.Create)
	router.DELETE("/api/v1/images/:image_id", imagesController.Delete)
	router.GET("/api/v1/images/:image_id", imagesController.FindById)

	// error handler
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server Run On Port:3000")
	err := server.ListenAndServe()
	helper.PanicError(err)
}
