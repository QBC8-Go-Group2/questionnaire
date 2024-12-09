package port

import (
	"context"

	"github.com/QBC8-Go-Group2/questionnaire/internal/auth/domain"
)

type Service interface {
	// Registration endpoints
	InitiateRegister(ctx context.Context, req domain.InitiateRegisterRequest) error
	CompleteRegister(ctx context.Context, req domain.CompleteRegisterRequest) error // Changed return type to just error

	// Login endpoints
	InitiateLogin(ctx context.Context, req domain.InitiateLoginRequest) error
	CompleteLogin(ctx context.Context, req domain.CompleteLoginRequest) (string, error)
}
