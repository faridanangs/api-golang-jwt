package images

import "mime/multipart"

type ImageRequest struct {
	FormData []*multipart.FileHeader `validate:"required"`
}
