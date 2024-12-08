package port

import (
	"context"

	"github.com/QBC8-Go-Group2/questionnaire/internal/auth/domain"
)

type Service interface {
	Register(ctx context.Context, req domain.RegisterRequest) (string, error)
	InitiateOTP(ctx context.Context, email string) error
	VerifyOTP(ctx context.Context, req domain.OTPRequest) (string, error)
}
