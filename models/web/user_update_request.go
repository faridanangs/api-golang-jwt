package web

type UserUpdateRequest struct {
	Id        string `json:"id" validate:"required"`
	FirstName string `validate:"required,min=1,max=25" json:"first_name"`
	LastName  string `validate:"required,min=1,max=50" json:"last_name"`
}
