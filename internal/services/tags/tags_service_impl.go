package services

import (
	converter "crud-go-api/internal/converter/request"
	converter2 "crud-go-api/internal/converter/response"
	constant2 "crud-go-api/libs/common/constant/error/main"
	constant "crud-go-api/libs/common/constant/logger"
	"crud-go-api/libs/common/exception"
	helper "crud-go-api/libs/common/helper/error"
	logger "crud-go-api/libs/common/logger/main"
	repository "crud-go-api/libs/data-layer/repository/ent/tags"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func InjectTagsServiceImpl(tagsRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagsRepository,
		Validate:       validate,
	}
}

func (service TagsServiceImpl) Create(ctx *gin.Context) (converter2.TagsOutputResponse, error) {
	createTagsRequest := converter.CreateTagRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.CreateTagsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return converter2.TagsOutputResponse{}, err
	}

	err = service.Validate.Struct(createTagsRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.CreateTagsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return converter2.TagsOutputResponse{}, err
	}

	tagCreateData, err := service.TagsRepository.Save(createTagsRequest, ctx)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.CreateTagsService, Message: err.Error()})
		return converter2.TagsOutputResponse{}, err
	}

	return converter2.TagsOutputResponse{Id: tagCreateData.UUID}, nil
}

func (service TagsServiceImpl) Update(ctx *gin.Context) (converter2.TagsOutputResponse, error) {
	updateTagsRequest := converter.UpdateTagRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.UpdateTagsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return converter2.TagsOutputResponse{}, err
	}

	tagId := ctx.Param("tagId")
	parsedUUID, err := helper.IsValidUUID(tagId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.UpdateTagsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return converter2.TagsOutputResponse{}, err
	}

	updateTagsRequest.Id = parsedUUID

	result, err := service.TagsRepository.Update(updateTagsRequest, ctx)
	if err != nil {
		return converter2.TagsOutputResponse{}, err
	}

	return converter2.TagsOutputResponse{Id: result}, nil
}

func (service TagsServiceImpl) Delete(ctx *gin.Context) (converter2.TagsResponse, error) {
	tagId := ctx.Param("tagId")
	uuidParse, err := helper.IsValidUUID(tagId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.DeleteTagsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return converter2.TagsResponse{}, err
	}

	deletedTagId, err := service.TagsRepository.Delete(uuidParse, ctx)
	if err != nil {
		return converter2.TagsResponse{}, err
	}

	return converter2.TagsResponse{Id: deletedTagId}, nil
}

func (service TagsServiceImpl) FindById(ctx *gin.Context) (converter2.TagsResponse, error) {
	tagId := ctx.Param("tagId")

	id, err := strconv.Atoi(tagId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.FindByIdTagsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return converter2.TagsResponse{}, err
	}

	tagData, err := service.TagsRepository.FindById(id, ctx)
	if err != nil {
		return converter2.TagsResponse{}, err
	}

	return converter2.TagsResponse{
		Id:   tagData.UUID,
		Name: tagData.Name,
	}, nil
}

func (service TagsServiceImpl) FindAll(ctx *gin.Context) ([]converter2.TagsResponse, error) {
	result, err := service.TagsRepository.FindAll()
	if err != nil {
		webError := exception.Error{
			Message: constant2.DBInternalError,
		}
		webError.InternalException(ctx)
		return nil, err
	}

	var tags []converter2.TagsResponse
	for _, value := range result {
		tag := converter2.TagsResponse{
			Id:   value.UUID,
			Name: value.Name,
		}

		tags = append(tags, tag)
	}
	return tags, nil
}
