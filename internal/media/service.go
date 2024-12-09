package media

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/QBC8-Go-Group2/questionnaire/internal/media/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/media/port"
	"github.com/google/uuid"
)

type service struct {
	repo       port.Repo
	uploadPath string
}

func NewService(repo port.Repo, uploadPath string) port.Service {
	return &service{
		repo:       repo,
		uploadPath: uploadPath,
	}
}

func (s *service) Upload(ctx context.Context, userID uint, file *multipart.FileHeader) (domain.Media, error) {
	err := os.MkdirAll(s.uploadPath, os.ModePerm)
	if err != nil {
		return domain.Media{}, fmt.Errorf("failed to create upload directory: %w", err)
	}

	mediaUUID := uuid.New().String()
	filename := fmt.Sprintf("%s_%s", mediaUUID, filepath.Base(file.Filename))
	filepath := filepath.Join(s.uploadPath, filename)

	src, err := file.Open()
	if err != nil {
		return domain.Media{}, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(filepath)
	if err != nil {
		return domain.Media{}, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return domain.Media{}, fmt.Errorf("failed to copy file content: %w", err)
	}

	media := domain.Media{
		UUID:      domain.MediaUUID(mediaUUID),
		UserID:    userID,
		Path:      filepath,
		Type:      domain.MediaType(file.Header.Get("Content-Type")),
		Size:      file.Size,
		Name:      file.Filename,
		CreatedAt: time.Now(),
	}

	id, err := s.repo.Create(ctx, media)
	if err != nil {
		return domain.Media{}, err
	}
	media.ID = id
	return media, nil
}

func (s *service) GetByID(ctx context.Context, id domain.MediaID) (domain.Media, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) GetByUserID(ctx context.Context, userID uint) ([]domain.Media, error) {
	return s.repo.FindByUserID(ctx, userID)
}
