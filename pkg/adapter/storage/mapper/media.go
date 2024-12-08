package mapper

import (
	"github.com/QBC8-Go-Group2/questionnaire/internal/media/domain"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
)

func MediaStorage2Domain(media types.Media) domain.Media {
	return domain.Media{
		ID:        domain.MediaID(media.ID),
		UserID:    media.UserID,
		Path:      media.Path,
		Type:      domain.MediaType(media.Type),
		Size:      media.Size,
		Name:      media.Name,
		CreatedAt: media.CreatedAt,
	}
}

func MediaDomain2Storage(media domain.Media) types.Media {
	return types.Media{
		ID:        uint(media.ID),
		UserID:    media.UserID,
		Path:      media.Path,
		Type:      string(media.Type),
		Size:      media.Size,
		Name:      media.Name,
		CreatedAt: media.CreatedAt,
	}
}
