package converter

import "github.com/google/uuid"

type CreateTagRequest struct {
	Name string `validate:"required,min=1,max=10" json:"name"`
}

type UpdateTagRequest struct {
	Id   uuid.UUID `validate:"required"`
	Name string    `validate:"required,max=200,min=1" json:"name"`
}
