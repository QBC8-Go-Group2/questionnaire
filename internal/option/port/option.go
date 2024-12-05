package port

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/option/domain"
)

type Repo interface {
	Create(ctx context.Context, option domain.Option) (domain.OptionID, error)
	Update(ctx context.Context, option domain.Option) (domain.OptionID, error)
	FindByID(ctx context.Context, optionID domain.OptionID) (domain.Option, error)
	DeleteWithID(ctx context.Context, option domain.OptionID) error
}
