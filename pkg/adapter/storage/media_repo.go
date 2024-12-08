package storage

import (
	"context"

	"github.com/QBC8-Go-Group2/questionnaire/internal/media/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/media/port"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/mapper"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
	"gorm.io/gorm"
)

type mediaRepo struct {
	db *gorm.DB
}

func NewMediaRepo(db *gorm.DB) port.Repo {
	return &mediaRepo{db: db}
}

func (r *mediaRepo) Create(ctx context.Context, media domain.Media) (domain.MediaID, error) {
	mediaStorage := mapper.MediaDomain2Storage(media)
	err := r.db.WithContext(ctx).Create(&mediaStorage).Error
	return domain.MediaID(mediaStorage.ID), err
}

func (r *mediaRepo) FindByID(ctx context.Context, id domain.MediaID) (domain.Media, error) {
	var media types.Media
	err := r.db.WithContext(ctx).First(&media, id).Error
	return mapper.MediaStorage2Domain(media), err
}

func (r *mediaRepo) FindByUserID(ctx context.Context, userID uint) ([]domain.Media, error) {
	var medias []types.Media
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&medias).Error
	if err != nil {
		return nil, err
	}

	result := make([]domain.Media, len(medias))
	for i, media := range medias {
		result[i] = mapper.MediaStorage2Domain(media)
	}
	return result, nil
}
