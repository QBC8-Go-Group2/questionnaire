package port

import (
	"context"

	"github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserDbID, error)
	Update(ctx context.Context, user domain.User) (domain.UserDbID, error)
	FindWithUserID(ctx context.Context, userId domain.UserID) (domain.User, error)
	FindWithUserDbID(ctx context.Context, userDbId domain.UserDbID) (domain.User, error)
	FindWithEmail(ctx context.Context, email string) (domain.User, error)
	DeleteWithUserID(ctx context.Context, user domain.UserID) error
	DeleteWithUserDbId(ctx context.Context, user domain.UserDbID) error
}
