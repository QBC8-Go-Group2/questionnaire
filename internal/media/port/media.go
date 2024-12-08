package port

import (
	"context"

	"github.com/QBC8-Go-Group2/questionnaire/internal/media/domain"
)

type Repo interface {
	Create(ctx context.Context, media domain.Media) (domain.MediaID, error)
	FindByID(ctx context.Context, id domain.MediaID) (domain.Media, error)
	FindByUserID(ctx context.Context, userID uint) ([]domain.Media, error)
}
