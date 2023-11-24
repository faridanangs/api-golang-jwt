package service

import (
	"context"
	"database/sql"
	"io"
	"os"
	"strings"
	"user_service/exception"
	"user_service/helper"
	"user_service/models/entity"
	"user_service/models/web/images"
	"user_service/repository"
	"user_service/utils"

	"github.com/go-playground/validator/v10"
)

type ImagesServiceIplm struct {
	DB               *sql.DB
	Validate         validator.Validate
	ImagesRepository repository.ImageRepository
}

func NewImagesServiceIplm(db *sql.DB, validate validator.Validate, imagesRepository repository.ImageRepository) ImagesService {
	return &ImagesServiceIplm{
		DB:               db,
		Validate:         validate,
		ImagesRepository: imagesRepository,
	}
}

func (service *ImagesServiceIplm) Create(ctx context.Context, request images.ImageRequest) []images.ImagesResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	var images []images.ImagesResponse

	for _, data := range request.FormData {
		file, err := data.Open()
		helper.PanicError(err)
		defer file.Close()

		tempFile, err := os.CreateTemp("public", "image-*.png")
		helper.PanicError(err)
		defer tempFile.Close()

		byteFileName, err := io.ReadAll(file)
		helper.PanicError(err)

		tempFile.Write(byteFileName)

		nameFile := tempFile.Name()

		newNamefile := strings.Split(nameFile, "\\")
		image := entity.Images{
			Id:        utils.GenerateId(),
			Path:      newNamefile[1],
			CreatedAt: int64(utils.GenerateTime()),
			UpdatedAt: int64(utils.GenerateTime()),
		}

		response := service.ImagesRepository.Save(ctx, tx, image)

		images = append(images, helper.ImagesResponse(response))
	}

	return images

}
func (service *ImagesServiceIplm) Delete(ctx context.Context, request images.ImagesResponse) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	response, err := service.ImagesRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ImagesRepository.Delete(ctx, tx, response)
	os.Remove("public" + response.Path)

}
func (service *ImagesServiceIplm) FindById(ctx context.Context, requestId string) images.ImagesResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	response, err := service.ImagesRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ImagesResponse(response)
}
