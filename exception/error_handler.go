package exception

import (
	"net/http"
	"user_service/helper"
	"user_service/models/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	if notFound(w, r, err) {
		return
	}
	if errStructValidator(w, r, err) {
		return
	}

	if unauthorized(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func notFound(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(ErrNotFound)

	if ok {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response := web.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}
		helper.WriteRequestToBody(w, response)
		return true
	} else {
		return false
	}

}
func unauthorized(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(ErrorUnauthorized)

	if ok {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		response := web.Response{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   exception.Error,
		}
		helper.WriteRequestToBody(w, response)
		return true
	} else {
		return false
	}

}
func errStructValidator(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := web.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}
		helper.WriteRequestToBody(w, response)
		return true
	} else {
		return false
	}

}

func internalServerError(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := web.Response{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteRequestToBody(w, response)
}
