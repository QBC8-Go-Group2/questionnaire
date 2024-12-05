package main

import (
	"log"

	"github.com/QBC8-Go-Group2/questionnaire/api/handler/http"
	"github.com/QBC8-Go-Group2/questionnaire/app"
	"github.com/QBC8-Go-Group2/questionnaire/config"
	"github.com/QBC8-Go-Group2/questionnaire/internal/auth"
	"github.com/QBC8-Go-Group2/questionnaire/internal/user"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/email"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.MustReadConfig("config.json")
	application := app.MustNewApp(cfg)

	emailService := email.NewService(email.Config{
		Host:     cfg.Email.Host,
		Port:     cfg.Email.Port,
		Username: cfg.Email.Username,
		Password: cfg.Email.Password,
	})

	userRepo := storage.NewUserRepo(application.DB())
	userService := user.NewService(userRepo)
	otpStore := storage.NewOTPStore(application.Redis())
	authService := auth.NewService(userService, otpStore, emailService)
	authHandler := http.NewAuthHandler(authService)

	app := fiber.New()
	http.RegisterAuthRoutes(app, authHandler)

	log.Fatal(app.Listen(":3000"))
}
