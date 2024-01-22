package services

import (
	converter "crud-go-api/internal/converter/response"
	"github.com/gin-gonic/gin"
)

type TagsService interface {
	Create(ctx *gin.Context) (converter.TagsOutputResponse, error)
	Update(ctx *gin.Context) (converter.TagsOutputResponse, error)
	Delete(ctx *gin.Context) (converter.TagsResponse, error)
	FindById(ctx *gin.Context) (converter.TagsResponse, error)
	FindAll(ctx *gin.Context) ([]converter.TagsResponse, error)
}
