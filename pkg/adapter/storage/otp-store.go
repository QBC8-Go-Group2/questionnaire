package storage

import (
	"context"
	"encoding/json"
	"time"

	"github.com/QBC8-Go-Group2/questionnaire/internal/auth/domain"
	"github.com/go-redis/redis/v8"
)

type otpStore struct {
	redis *redis.Client
}

func NewOTPStore(redis *redis.Client) *otpStore {
	return &otpStore{redis: redis}
}

func (s *otpStore) StoreOTP(ctx context.Context, data domain.OTPData) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	expiration := time.Until(data.ExpiresAt)
	return s.redis.Set(ctx, "otp:"+data.Email, bytes, expiration).Err()
}

func (s *otpStore) GetOTP(ctx context.Context, email string) (domain.OTPData, error) {
	var data domain.OTPData
	bytes, err := s.redis.Get(ctx, "otp:"+email).Bytes()
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(bytes, &data)
	return data, err
}

func (s *otpStore) DeleteOTP(ctx context.Context, email string) error {
	return s.redis.Del(ctx, "otp:"+email).Err()
}
