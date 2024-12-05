package response

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/response/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/response/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo}
}

func (s *service) CreateResponse(ctx context.Context, response domain.Response) (domain.ResponseID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateResponse(ctx context.Context, response domain.Response) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindResponseById(ctx context.Context, id domain.ResponseID) (domain.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteResponse(ctx context.Context, response domain.Response) error {
	//TODO implement me
	panic("implement me")
}
