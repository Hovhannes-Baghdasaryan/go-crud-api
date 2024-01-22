package converter

import "github.com/google/uuid"

type TagsResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name,omitempty"`
}

type TagsOutputResponse struct {
	Id uuid.UUID `json:"id"`
}
