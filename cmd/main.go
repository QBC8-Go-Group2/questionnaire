package main

import (
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
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

	questionnaireRepo := storage.NewQuestionnaireRepo(application.DB())
	questionnaireService := questionnaire.NewService(questionnaireRepo)
	questionnaireHandler := http.NewQuestionnaireHandler(questionnaireService)

	fiberApp := fiber.New()
	fiberApp.Use(logger.New())
	fiberApp.Use(recover2.New())
	fiberApp.Use(http.Limiter())

	baseGroup := fiberApp.Group("/api/v1")

	transaction := http.QuestionsTransaction(application.DB())
	http.RegisterAuthRoutes(baseGroup, authHandler)
	http.RegisterQuestionnaireRoutes(baseGroup, transaction, questionnaireHandler)

	log.Fatal(fiberApp.Listen(":3000"))
}
