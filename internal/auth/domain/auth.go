package domain

import "time"

type (
	// First step of registration
	InitiateRegisterRequest struct {
		Email string
	}

	// Second step of registration after OTP verification
	CompleteRegisterRequest struct {
		Email    string
		OTP      string
		Password string
		NatId    string
	}

	// First step of login
	InitiateLoginRequest struct {
		Email    string
		Password string
	}

	// Second step of login
	CompleteLoginRequest struct {
		Email string
		OTP   string
	}

	OTPRequest struct {
		Email string
		OTP   string
	}

	OTPData struct {
		Email     string
		OTP       string
		ExpiresAt time.Time
		Purpose   OTPPurpose // login, registration
	}

	OTPPurpose string
)

const (
	OTPPurposeLogin        OTPPurpose = "login"
	OTPPurposeRegistration OTPPurpose = "registration"
)
