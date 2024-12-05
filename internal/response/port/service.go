package port

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/response/domain"
)

type Service interface {
	CreateResponse(ctx context.Context, response domain.Response) (domain.ResponseID, error)
	UpdateResponse(ctx context.Context, response domain.Response) error
	FindResponseById(ctx context.Context, id domain.ResponseID) (domain.Response, error)
	DeleteResponse(ctx context.Context, response domain.Response) error
}
