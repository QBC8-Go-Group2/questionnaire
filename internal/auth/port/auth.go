package port

import (
	"context"

	"github.com/QBC8-Go-Group2/questionnaire/internal/auth/domain"
)

type OTPStore interface {
	StoreOTP(ctx context.Context, data domain.OTPData) error
	GetOTP(ctx context.Context, email string) (domain.OTPData, error)
	DeleteOTP(ctx context.Context, email string) error
}
