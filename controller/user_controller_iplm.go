package controller

import (
	"net/http"
	"user_service/helper"
	"user_service/models/web"
	"user_service/service"

	"github.com/julienschmidt/httprouter"
)

type UserControllerIplm struct {
	UserService service.UserService
}

func NewUserControllerIplm(userService service.UserService) UserController {
	return &UserControllerIplm{
		UserService: userService,
	}
}

func (controller *UserControllerIplm) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadRequestToBody(r, &userCreateRequest)

	response := controller.UserService.Create(r.Context(), userCreateRequest)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteRequestToBody(w, webResponse)

}
func (controller *UserControllerIplm) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userUpdateRequest := &web.UserUpdateRequest{}
	helper.ReadRequestToBody(r, userUpdateRequest)
	userUpdateRequest.Id = params.ByName("user_id")
	response := controller.UserService.Update(r.Context(), *userUpdateRequest)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteRequestToBody(w, webResponse)
}
func (controller *UserControllerIplm) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("user_id")
	controller.UserService.Delete(r.Context(), userId)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
	}

	helper.WriteRequestToBody(w, webResponse)
}
func (controller *UserControllerIplm) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("user_id")
	response := controller.UserService.FindById(r.Context(), userId)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteRequestToBody(w, webResponse)
}
func (controller *UserControllerIplm) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := controller.UserService.FindAll(r.Context())

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteRequestToBody(w, webResponse)
}
func (controller *UserControllerIplm) Auth(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userAuthRequest := &web.UserAuthRequest{}
	helper.ReadRequestToBody(r, userAuthRequest)

	response := controller.UserService.Auth(r.Context(), *userAuthRequest)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteRequestToBody(w, webResponse)
}

func (controller *UserControllerIplm) WithRefreshToken(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userToken := r.Header.Get("Authorization")

	response := controller.UserService.WithRefreshToken(r.Context(), userToken)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteRequestToBody(w, webResponse)
}
