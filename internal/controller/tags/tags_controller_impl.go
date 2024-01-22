package controller

import (
	services "crud-go-api/internal/services/tags"
	"crud-go-api/libs/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TagsControllerImpl struct {
	tagsService services.TagsService
}

func InjectTagsController(service services.TagsService) TagsController {
	return &TagsControllerImpl{
		tagsService: service,
	}
}

func (controller TagsControllerImpl) Create(ctx *gin.Context) {
	tagCreateData, err := controller.tagsService.Create(ctx)
	if err != nil {
		return
	}

	webResponse := response.Response{
		Status:  http.StatusCreated,
		Message: "Tag Created Successfully",
		Data:    tagCreateData,
	}
	webResponse.ActionSucceeded(ctx)
}

func (controller TagsControllerImpl) Update(ctx *gin.Context) {
	updatedTagId, err := controller.tagsService.Update(ctx)
	if err != nil {
		return
	}

	webResponse := response.Response{
		Message: "Updated Successfully",
		Data:    updatedTagId,
	}
	webResponse.ActionSucceeded(ctx)
}

func (controller TagsControllerImpl) Delete(ctx *gin.Context) {
	resData, err := controller.tagsService.Delete(ctx)
	if err != nil {
		return
	}

	webResponse := response.Response{
		Message: "Deleted successfully",
		Data:    resData,
	}
	webResponse.ActionSucceeded(ctx)
}

func (controller TagsControllerImpl) FindById(ctx *gin.Context) {
	tagResponse, err := controller.tagsService.FindById(ctx)
	if err != nil {
		return
	}

	webResponse := response.Response{
		Message: "Tag Single Find",
		Data:    tagResponse,
	}
	webResponse.ActionSucceeded(ctx)
}

func (controller TagsControllerImpl) FindAll(ctx *gin.Context) {
	tagResponseAll, err := controller.tagsService.FindAll(ctx)
	if err != nil {
		return
	}

	webResponse := response.Response{
		Message: "Tag All Find",
		Data:    tagResponseAll,
	}
	webResponse.ActionSucceeded(ctx)
}
