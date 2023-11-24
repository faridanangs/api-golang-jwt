package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ImagesController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
