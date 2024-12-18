package port

import (
	"context"
	"mime/multipart"

	"github.com/QBC8-Go-Group2/questionnaire/internal/media/domain"
)

type Service interface {
	Upload(ctx context.Context, userID uint, file *multipart.FileHeader) (domain.Media, error)
	GetByID(ctx context.Context, id domain.MediaID) (domain.Media, error)
	GetByUUID(ctx context.Context, uuid domain.MediaUUID) (domain.Media, error)
	GetByUserID(ctx context.Context, userID uint) ([]domain.Media, error)
	VerifyFileAccess(ctx context.Context, mediaUUID domain.MediaUUID, userID uint) (domain.Media, error)
}
