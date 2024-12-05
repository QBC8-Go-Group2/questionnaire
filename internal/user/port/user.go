package port

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserDbID, error)
	Update(ctx context.Context, user domain.User) (domain.UserDbID, error)
	FindWithUserDbID(ctx context.Context, userDbId domain.UserDbID) (domain.User, error)
	DeleteWithUserDbId(ctx context.Context, user domain.UserDbID) error
}
