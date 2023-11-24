package service

import (
	"context"
	"user_service/models/web/images"
)

type ImagesService interface {
	Create(ctx context.Context, request images.ImageRequest) []images.ImagesResponse
	Delete(ctx context.Context, request images.ImagesResponse)
	FindById(ctx context.Context, requestId string) images.ImagesResponse
}
