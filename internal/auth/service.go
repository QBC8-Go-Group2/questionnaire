package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	mathrand "math/rand"
	"net/mail"
	"time"

	"github.com/QBC8-Go-Group2/questionnaire/internal/auth/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/auth/port"
	userDomain "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	userPort "github.com/QBC8-Go-Group2/questionnaire/internal/user/port"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/email"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userService  userPort.Service
	otpStore     port.OTPStore
	emailService email.Service
	jwtService   jwt.JWTGenerator
}

func NewService(userService userPort.Service, otpStore port.OTPStore, emailService email.Service, jwtService jwt.JWTGenerator) port.Service {
	return &service{
		userService:  userService,
		otpStore:     otpStore,
		emailService: emailService,
		jwtService:   jwtService,
	}
}

// InitiateRegister starts the registration process by sending OTP
func (s *service) InitiateRegister(ctx context.Context, req domain.InitiateRegisterRequest) error {
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return fmt.Errorf("invalid email format")
	}

	// Check if user already exists
	_, err := s.userService.FindUserWithEmail(ctx, req.Email)
	if err == nil {
		return fmt.Errorf("user already exists")
	}

	otp := generateOTP()
	err = s.emailService.SendOTP(req.Email, otp)
	if err != nil {
		return fmt.Errorf("failed to send OTP: %w", err)
	}

	err = s.otpStore.StoreOTP(ctx, domain.OTPData{
		Email:     req.Email,
		OTP:       otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
		Purpose:   domain.OTPPurposeRegistration,
	})
	if err != nil {
		return fmt.Errorf("failed to store OTP: %w", err)
	}

	return nil
}

// CompleteRegister completes the registration after OTP verification
func (s *service) CompleteRegister(ctx context.Context, req domain.CompleteRegisterRequest) error {
	// Verify OTP
	otpData, err := s.otpStore.GetOTP(ctx, req.Email)
	if err != nil {
		return fmt.Errorf("invalid OTP request: %w", err)
	}

	if otpData.Purpose != domain.OTPPurposeRegistration {
		return fmt.Errorf("invalid OTP purpose")
	}

	if time.Now().After(otpData.ExpiresAt) {
		return fmt.Errorf("OTP expired")
	}

	if otpData.OTP != req.OTP {
		return fmt.Errorf("invalid OTP")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
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
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Clean up OTP
	_ = s.otpStore.DeleteOTP(ctx, req.Email)

	return nil
}

// InitiateLogin starts the login process by verifying password and sending OTP
func (s *service) InitiateLogin(ctx context.Context, req domain.InitiateLoginRequest) error {
	user, err := s.userService.FindUserWithEmail(ctx, req.Email)
	if err != nil {
		return fmt.Errorf("invalid credentials")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return fmt.Errorf("invalid credentials")
	}

	// Generate and send OTP
	otp := generateOTP()
	err = s.emailService.SendOTP(req.Email, otp)
	if err != nil {
		return fmt.Errorf("failed to send OTP: %w", err)
	}

	err = s.otpStore.StoreOTP(ctx, domain.OTPData{
		Email:     req.Email,
		OTP:       otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
		Purpose:   domain.OTPPurposeLogin,
	})
	if err != nil {
		return fmt.Errorf("failed to store OTP: %w", err)
	}

	return nil
}

// CompleteLogin completes the login process by verifying OTP
func (s *service) CompleteLogin(ctx context.Context, req domain.CompleteLoginRequest) (string, error) {
	otpData, err := s.otpStore.GetOTP(ctx, req.Email)
	if err != nil {
		return "", fmt.Errorf("invalid OTP request: %w", err)
	}

	if otpData.Purpose != domain.OTPPurposeLogin {
		return "", fmt.Errorf("invalid OTP purpose")
	}

	if time.Now().After(otpData.ExpiresAt) {
		return "", fmt.Errorf("OTP expired")
	}

	if otpData.OTP != req.OTP {
		return "", fmt.Errorf("invalid OTP")
	}

	// Get user for JWT generation
	user, err := s.userService.FindUserWithEmail(ctx, req.Email)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err)
	}

	// Clean up OTP
	_ = s.otpStore.DeleteOTP(ctx, req.Email)

	// Generate JWT
	token, err := s.jwtService.GenerateJWT(string(user.Role), user.Email, uint(user.ID))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

// Helper functions
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
