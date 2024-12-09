package port

import (
	"context"

	"github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserDbID, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.UserDbID, error)
	FindUserWithUserDbID(ctx context.Context, userDbId domain.UserDbID) (domain.User, error)
	FindUserWithEmail(ctx context.Context, email string) (domain.User, error)
	DeleteUserWithUserID(ctx context.Context, user domain.UserID) error
	DeleteUserWithUserDbId(ctx context.Context, user domain.UserDbID) error
}
