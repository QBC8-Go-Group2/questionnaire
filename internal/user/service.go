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
	return s.repo.Create(ctx, user)
}

func (s *service) UpdateUser(ctx context.Context, user domain.User) (domain.UserDbID, error) {
	return s.repo.Update(ctx, user)
}

func (s *service) FindUserWithUserID(ctx context.Context, userId domain.UserID) (domain.User, error) {
	return s.repo.FindWithUserID(ctx, userId)
}

func (s *service) FindUserWithUserDbID(ctx context.Context, userDbId domain.UserDbID) (domain.User, error) {
	return s.repo.FindWithUserDbID(ctx, userDbId)
}

func (s *service) FindUserWithEmail(ctx context.Context, email string) (domain.User, error) {
	return s.repo.FindWithEmail(ctx, email)
}

func (s *service) DeleteUserWithUserID(ctx context.Context, user domain.UserID) error {
	return s.repo.DeleteWithUserID(ctx, user)
}

func (s *service) DeleteUserWithUserDbId(ctx context.Context, user domain.UserDbID) error {
	return s.repo.DeleteWithUserDbId(ctx, user)
}
