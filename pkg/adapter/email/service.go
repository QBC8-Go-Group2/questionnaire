package email

import (
	"fmt"
	"net/smtp"
)

type Service interface {
	SendOTP(to, otp string) error
}

type service struct {
	config Config
}

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewService(cfg Config) Service {
	return &service{config: cfg}
}

func (s *service) SendOTP(to, otp string) error {
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)

	msg := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: Your OTP Code\r\n\r\n"+
		"Your OTP code is: %s\r\n", s.config.Username, to, otp))

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
		auth,
		s.config.Username,
		[]string{to},
		msg,
	)
}
