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

func (s *service) Upload(ctx context.Context, userID uint, file *multipart.FileHeader) (domain.MediaID, error) {
	// Create upload directory if it doesn't exist
	err := os.MkdirAll(s.uploadPath, os.ModePerm)
	if err != nil {
		return 0, fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
	filepath := filepath.Join(s.uploadPath, filename)

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return 0, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	// Create destination file
	dst, err := os.Create(filepath)
	if err != nil {
		return 0, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err = io.Copy(dst, src); err != nil {
		return 0, fmt.Errorf("failed to copy file content: %w", err)
	}

	// Create media record
	media := domain.Media{
		UserID:    userID,
		Path:      filepath,
		Type:      domain.MediaType(file.Header.Get("Content-Type")),
		Size:      file.Size,
		Name:      file.Filename,
		CreatedAt: time.Now(),
	}

	return s.repo.Create(ctx, media)
}

func (s *service) GetByID(ctx context.Context, id domain.MediaID) (domain.Media, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) GetByUserID(ctx context.Context, userID uint) ([]domain.Media, error) {
	return s.repo.FindByUserID(ctx, userID)
}
