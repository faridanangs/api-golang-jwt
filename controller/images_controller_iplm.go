package controller

import (
	"net/http"
	"os"
	"user_service/helper"
	"user_service/models/web"
	"user_service/models/web/images"
	"user_service/service"

	"github.com/julienschmidt/httprouter"
)

type ImagesControllerIplm struct {
	ImagesService service.ImagesService
}

func NewImagesControllerIplm(imagesService service.ImagesService) ImagesController {
	return &ImagesControllerIplm{
		ImagesService: imagesService,
	}
}

func (controller *ImagesControllerIplm) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseMultipartForm(10 * 1024 * 1024)

	imageResuest := images.ImageRequest{}
	imageResuest.FormData = r.MultipartForm.File["image"]

	response := controller.ImagesService.Create(r.Context(), imageResuest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}

	helper.WriteRequestToBody(w, webResponse)

}

func (controller *ImagesControllerIplm) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	imageId := params.ByName("image_id")

	respoonse := controller.ImagesService.FindById(r.Context(), imageId)
	controller.ImagesService.Delete(r.Context(), respoonse)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
	}
	helper.WriteRequestToBody(w, webResponse)
}
func (controller *ImagesControllerIplm) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	imageId := params.ByName("image_id")

	response := controller.ImagesService.FindById(r.Context(), imageId)
	byteFile, err := os.ReadFile("public/" + response.Path)
	helper.PanicError(err)
	w.Write(byteFile)
}
