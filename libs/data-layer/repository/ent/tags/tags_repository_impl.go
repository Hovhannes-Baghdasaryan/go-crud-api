package repository

import (
	"context"
	converter "crud-go-api/internal/converter/request"
	constant2 "crud-go-api/libs/common/constant/error/main"
	constant "crud-go-api/libs/common/constant/error/tag"
	constant3 "crud-go-api/libs/common/constant/logger"
	"crud-go-api/libs/common/exception"
	logger "crud-go-api/libs/common/logger/main"
	"crud-go-api/libs/data-layer/entity/tags/ent"
	"crud-go-api/libs/data-layer/entity/tags/ent/tagentity"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type TagsRepositoryImpl struct {
	clientDB *ent.Client
}

func InjectTagsRepositoryImpl(clientDB *ent.Client) *TagsRepositoryImpl {
	return &TagsRepositoryImpl{
		clientDB: clientDB,
	}
}

func (repo *TagsRepositoryImpl) Save(tag converter.CreateTagRequest, ctx *gin.Context) (*ent.TagEntity, error) {
	result, err := repo.clientDB.TagEntity.Create().SetName(tag.Name).Save(context.Background())
	if err != nil {
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return &ent.TagEntity{}, err
	}

	return result, nil
}

func (repo *TagsRepositoryImpl) Update(updatePayload converter.UpdateTagRequest, ctx *gin.Context) (uuid.UUID, error) {
	ok, err := repo.clientDB.TagEntity.Update().SetName(updatePayload.Name).Where(tagentity.UUID(updatePayload.Id)).Save(context.Background())

	if ok == 0 {
		logger.LogError(logger.LoggerPayload{FuncName: constant3.UpdateTagsRepository, Message: constant.NotFound})
		webError := exception.Error{
			Message: errors.New(constant.NotFound).Error(),
		}
		webError.NotFoundException(ctx)
		return uuid.UUID{}, errors.New(constant.NotFound)
	}

	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant3.UpdateTagsRepository, Message: constant2.DBInternalError})
		webError := exception.Error{
			Message: errors.New(constant2.DBInternalError).Error(),
		}
		webError.InternalException(ctx)
		return uuid.UUID{}, errors.New(constant2.DBInternalError)
	}

	return updatePayload.Id, nil
}

func (repo *TagsRepositoryImpl) Delete(tagUUId uuid.UUID, ctx *gin.Context) (uuid.UUID, error) {
	ok, err := repo.clientDB.TagEntity.Delete().Where(tagentity.UUID(tagUUId)).Exec(context.Background())
	if ok == 0 {
		logger.LogError(logger.LoggerPayload{FuncName: constant3.DeleteTagsRepository, Message: constant.NotFound})
		webError := exception.Error{
			Message: errors.New(constant.NotFound).Error(),
		}
		webError.NotFoundException(ctx)
		return uuid.UUID{}, errors.New(constant.NotFound)
	}

	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant3.DeleteTagsRepository, Message: constant2.DBInternalError})
		webError := exception.Error{
			Message: errors.New(constant2.DBInternalError).Error(),
		}
		webError.InternalException(ctx)
		return uuid.UUID{}, errors.New(constant2.DBInternalError)
	}

	return tagUUId, nil
}

func (repo *TagsRepositoryImpl) FindById(tagId int, ctx *gin.Context) (*ent.TagEntity, error) {
	result, err := repo.clientDB.TagEntity.Query().Where(tagentity.ID(tagId)).Only(context.Background())
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant3.FindByIdTagsRepository, Message: constant.NotFound})
		webError := exception.Error{
			Message: errors.New(constant.NotFound).Error(),
		}
		webError.NotFoundException(ctx)
		return nil, err
	}

	return result, nil
}

func (repo *TagsRepositoryImpl) FindAll() ([]*ent.TagEntity, error) {
	result, err := repo.clientDB.TagEntity.Query().All(context.Background())

	if err != nil {
		return nil, err
	}

	return result, nil
}
