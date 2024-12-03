package port

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/response/domain"
)

type Repo interface {
	Create(ctx context.Context, response domain.Response) (domain.ResponseID, error)
	Update(ctx context.Context, response domain.Response) error
	FindById(ctx context.Context, id domain.ResponseID) (domain.Response, error)
	Delete(ctx context.Context, response domain.Response) error
}
