package main

import (
	"log"
	"path/filepath"

	"github.com/QBC8-Go-Group2/questionnaire/api/handler/http"
	"github.com/QBC8-Go-Group2/questionnaire/app"
	"github.com/QBC8-Go-Group2/questionnaire/config"
	"github.com/QBC8-Go-Group2/questionnaire/internal/auth"
	"github.com/QBC8-Go-Group2/questionnaire/internal/media"
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire"
	"github.com/QBC8-Go-Group2/questionnaire/internal/user"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/email"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Initialize configuration and application
	cfg := config.MustReadConfig("config.json")
	application := app.MustNewApp(cfg)

	// Initialize services
	emailService := email.NewService(email.Config{
		Host:     cfg.Email.Host,
		Port:     cfg.Email.Port,
		Username: cfg.Email.Username,
		Password: cfg.Email.Password,
	})

	// Initialize JWT service
	jwtService := &jwt.JWTService{}

	// Initialize repositories
	userRepo := storage.NewUserRepo(application.DB())
	mediaRepo := storage.NewMediaRepo(application.DB())
	otpStore := storage.NewOTPStore(application.Redis())
	questionnaireRepo := storage.NewQuestionnaireRepo(application.DB())

	// Initialize domain services
	userService := user.NewService(userRepo)
	authService := auth.NewService(userService, otpStore, emailService, jwtService)
	questionnaireService := questionnaire.NewService(questionnaireRepo)

	// Initialize media service with upload path
	uploadPath := filepath.Join("pkg", "data", "uploads")
	mediaService := media.NewService(mediaRepo, uploadPath)

	// Initialize HTTP handlers
	authHandler := http.NewAuthHandler(authService)
	mediaHandler := http.NewMediaHandler(mediaService)
	questionnaireHandler := http.NewQuestionnaireHandler(questionnaireService)

	// Load public key for JWT
	pubKey, err := http.LoadPublicKey("Public_key.pem")
	if err != nil {
		log.Fatalf("Failed to load public key: %v", err)
	}
	log.Printf("Public key loaded successfully: %v", pubKey != nil)

	// Initialize single Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(http.Limiter())

	// Create base API group
	api := app.Group("/api/v1")

	// Register routes
	http.RegisterAuthRoutes(api, authHandler)
	http.RegisterMediaRoutes(api, mediaHandler)

	// Protected routes
	protected := api.Group("/protected")
	protected.Use(http.JWTMiddleware())

	// Profile route
	protected.Get("/profile", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id": c.Locals(http.UserIDKey),
			"email":   c.Locals(http.EmailKey),
			"role":    c.Locals(http.RoleKey),
		})
	})

	// Health check route
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Questionnaire routes with transaction
	transaction := http.QuestionsTransaction(application.DB())
	http.RegisterQuestionnaireRoutes(api, transaction, questionnaireHandler)

	// Start server
	log.Printf("Server starting on port :3000")
	log.Fatal(app.Listen(":3000"))
}
