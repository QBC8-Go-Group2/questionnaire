package storage

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/option/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/option/port"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/mapper"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
	"gorm.io/gorm"
)

type optionRepo struct {
	db *gorm.DB
}

func NewOptionRepo(db *gorm.DB) port.Repo {
	return &optionRepo{
		db: db,
	}
}

func (o *optionRepo) Create(ctx context.Context, option domain.Option) (domain.OptionID, error) {
	oStorage := mapper.OptionDomain2Storage(option)
	return domain.OptionID(oStorage.ID), o.db.Table("options").WithContext(ctx).Create(&oStorage).Error
}

func (o *optionRepo) Update(ctx context.Context, option domain.Option) (domain.OptionID, error) {
	oStorage := mapper.OptionDomain2Storage(option)
	return domain.OptionID(oStorage.ID), o.db.Table("options").WithContext(ctx).Updates(&oStorage).Error
}

func (o *optionRepo) FindByID(ctx context.Context, optionID domain.OptionID) (domain.Option, error) {
	var opt types.Option
	err := o.db.Table("options").WithContext(ctx).Where("id = ?", optionID).First(&opt).Error
	if err != nil {
		return domain.Option{}, err
	}
	return mapper.OptionStorage2Domain(opt), nil
}

func (o *optionRepo) DeleteWithID(ctx context.Context, option domain.OptionID) error {
	return o.db.Table("options").WithContext(ctx).Where("id = ?", option).Delete(&types.Option{}).Error
}
