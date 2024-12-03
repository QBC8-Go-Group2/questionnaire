package user

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/user/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo}
}

func (s *service) CreateUser(ctx context.Context, user domain.User) (domain.UserDbID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateUser(ctx context.Context, user domain.User) (domain.UserDbID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindUserWithUserID(ctx context.Context, userId domain.UserID) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindUserWithUserDbID(ctx context.Context, userDbId domain.UserDbID) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteUserWithUserID(ctx context.Context, user domain.UserID) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteUserWithUserDbId(ctx context.Context, user domain.UserDbID) error {
	//TODO implement me
	panic("implement me")
}
