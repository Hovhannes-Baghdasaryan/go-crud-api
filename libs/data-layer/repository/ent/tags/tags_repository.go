package repository

import (
	converter "crud-go-api/internal/converter/request"
	"crud-go-api/libs/data-layer/entity/tags/ent"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TagsRepository interface {
	Save(tags converter.CreateTagRequest, ctx *gin.Context) (*ent.TagEntity, error)
	Update(updatePayload converter.UpdateTagRequest, ctx *gin.Context) (uuid.UUID, error)
	Delete(tagId uuid.UUID, ctx *gin.Context) (uuid.UUID, error)
	FindById(tagId int, ctx *gin.Context) (*ent.TagEntity, error)
	FindAll() ([]*ent.TagEntity, error)
}
