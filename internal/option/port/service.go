package port

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/option/domain"
)

type Service interface {
	CreateOption(ctx context.Context, option domain.Option) (domain.OptionID, error)
	UpdateOption(ctx context.Context, option domain.Option) (domain.OptionID, error)
	FindOptionByID(ctx context.Context, optionID domain.OptionID) (domain.Option, error)
	DeleteOptionWithID(ctx context.Context, option domain.OptionID) error
}
