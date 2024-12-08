package domain

import "time"

type (
	LoginRequest struct {
		Email    string
		Password string
	}

	RegisterRequest struct {
		Email    string
		Password string
		NatId    string
	}

	OTPRequest struct {
		Email string
		OTP   string
	}

	OTPData struct {
		Email     string
		OTP       string
		ExpiresAt time.Time
	}
)
