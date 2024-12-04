package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	mathrand "math/rand"
	"time"

	"github.com/QBC8-Go-Group2/questionnaire/internal/auth/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/auth/port"
	userDomain "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	userPort "github.com/QBC8-Go-Group2/questionnaire/internal/user/port"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userService userPort.Service
	otpStore    port.OTPStore
}

func NewService(userService userPort.Service, otpStore port.OTPStore) port.Service {
	return &service{
		userService: userService,
		otpStore:    otpStore,
	}
}

func (s *service) Register(ctx context.Context, req domain.RegisterRequest) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	userID := generateUserID()
	user := userDomain.User{
		UserID:    userDomain.UserID(userID),
		Email:     req.Email,
		Password:  string(hashedPassword),
		NatId:     req.NatId,
		Role:      userDomain.UserRole,
		CreatedAT: time.Now(),
	}

	_, err = s.userService.CreateUser(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return userID, nil
}

func (s *service) Login(ctx context.Context, req domain.LoginRequest) (string, error) {
	return "", fmt.Errorf("direct login disabled, use OTP flow")
}

func (s *service) InitiateOTP(ctx context.Context, email string) error {
	otp := generateOTP()
	fmt.Printf("Generated OTP: %s\n", otp)

	err := s.otpStore.StoreOTP(ctx, domain.OTPData{
		Email:     email,
		OTP:       otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	})
	if err != nil {
		return fmt.Errorf("failed to store OTP: %w", err)
	}

	return nil
}

func (s *service) VerifyOTP(ctx context.Context, req domain.OTPRequest) (string, error) {
	otpData, err := s.otpStore.GetOTP(ctx, req.Email)
	if err != nil {
		return "", fmt.Errorf("invalid OTP request: %w", err)
	}

	if time.Now().After(otpData.ExpiresAt) {
		return "", fmt.Errorf("OTP expired")
	}

	if otpData.OTP != req.OTP {
		return "", fmt.Errorf("invalid OTP")
	}

	token := "dummy-jwt-token"

	err = s.otpStore.DeleteOTP(ctx, req.Email)
	if err != nil {
		return "", fmt.Errorf("failed to delete OTP: %w", err)
	}

	return token, nil
}

func generateUserID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func generateOTP() string {
	const digits = "0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = digits[mathrand.Intn(len(digits))]
	}
	return string(b)
}
