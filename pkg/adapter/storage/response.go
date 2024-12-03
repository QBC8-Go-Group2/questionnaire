package storage

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/response/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/response/port"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/mapper"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
	"gorm.io/gorm"
)

type responseRepo struct {
	db *gorm.DB
}

func NewResponseRepo(db *gorm.DB) port.Repo {
	return &responseRepo{
		db: db,
	}
}

func (r *responseRepo) Create(ctx context.Context, response domain.Response) (domain.ResponseID, error) {
	responseStorage := mapper.ResponseDomain2Storage(response)
	return domain.ResponseID(responseStorage.ID), r.db.Table("responses").WithContext(ctx).Create(&responseStorage).Error
}

func (r *responseRepo) Update(ctx context.Context, response domain.Response) error {
	responseStorage := mapper.ResponseDomain2Storage(response)
	return r.db.Table("responses").WithContext(ctx).Updates(&responseStorage).Error
}

func (r *responseRepo) FindById(ctx context.Context, id domain.ResponseID) (domain.Response, error) {
	var response types.Response
	err := r.db.Table("responses").Where("id = ?", id).First(&response).Error
	if err != nil {
		return domain.Response{}, err
	}
	return mapper.ResponseStorage2Domain(response), nil
}

func (r *responseRepo) Delete(ctx context.Context, response domain.Response) error {
	return r.db.Table("responses").Where("id = ?", response.ID).Delete(&response).Error
}
