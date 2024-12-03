package option

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/option/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/option/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo}
}

func (s *service) CreateOption(ctx context.Context, option domain.Option) (domain.OptionID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateOption(ctx context.Context, option domain.Option) (domain.OptionID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindOptionByID(ctx context.Context, optionID domain.OptionID) (domain.Option, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteOptionWithID(ctx context.Context, option domain.OptionID) error {
	//TODO implement me
	panic("implement me")
}
