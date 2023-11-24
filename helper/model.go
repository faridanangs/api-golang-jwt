package helper

import (
	"user_service/models/entity"
	"user_service/models/web"
	"user_service/models/web/images"
)

func UserResponse(user entity.Users) web.UsersResponse {
	return web.UsersResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func UserResponses(users []entity.Users) []web.UsersResponse {
	var responses []web.UsersResponse
	for _, data := range users {
		responses = append(responses, UserResponse(data))
	}
	return responses
}

func ImagesResponse(request entity.Images) images.ImagesResponse {
	return images.ImagesResponse{
		Id:        request.Id,
		Path:      request.Path,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}
}
