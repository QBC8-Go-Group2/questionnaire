package main

import (
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire"
	"log"
	"path/filepath"

	"github.com/QBC8-Go-Group2/questionnaire/api/handler/http"
	"github.com/QBC8-Go-Group2/questionnaire/app"
	"github.com/QBC8-Go-Group2/questionnaire/config"
	"github.com/QBC8-Go-Group2/questionnaire/internal/auth"
	"github.com/QBC8-Go-Group2/questionnaire/internal/media"
	"github.com/QBC8-Go-Group2/questionnaire/internal/user"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/email"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	// Initialize JWT service
	jwtService := &jwt.JWTService{}

	// Initialize repositories
	userRepo := storage.NewUserRepo(application.DB())
	mediaRepo := storage.NewMediaRepo(application.DB())
	otpStore := storage.NewOTPStore(application.Redis())

	// Initialize services
	userService := user.NewService(userRepo)
	authService := auth.NewService(userService, otpStore, emailService, jwtService)

	// Initialize media service with upload path
	uploadPath := filepath.Join("pkg", "data", "uploads")
	mediaService := media.NewService(mediaRepo, uploadPath)

	// Initialize HTTP handlers
	authHandler := http.NewAuthHandler(authService)
	mediaHandler := http.NewMediaHandler(mediaService)

	// Initialize Fiber app
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

	// Load and initialize the JWT middleware
	pubKey, err := http.LoadPublicKey("Public_key.pem")
	if err != nil {
		log.Fatalf("Failed to load public key: %v", err)
	}
	log.Printf("Public key loaded successfully: %v", pubKey != nil)

	// Add basic middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Register routes
	api := app.Group("/api/v1")

	// Auth routes
	http.RegisterAuthRoutes(app, authHandler)
	questionnaireRepo := storage.NewQuestionnaireRepo(application.DB())
	questionnaireService := questionnaire.NewService(questionnaireRepo)
	questionnaireHandler := http.NewQuestionnaireHandler(questionnaireService)

	fiberApp := fiber.New()
	fiberApp.Use(logger.New())
	fiberApp.Use(recover.New())
	fiberApp.Use(http.Limiter())

	// Protected routes using the middleware from middlewares.go
	protected := api.Group("/protected")
	protected.Use(http.JWTMiddleware())

	// Media routes
	http.RegisterMediaRoutes(app, mediaHandler)

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
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Start server
	port := ":3000"
	log.Printf("Server starting on port %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	baseGroup := fiberApp.Group("/api/v1")

	transaction := http.QuestionsTransaction(application.DB())
	http.RegisterAuthRoutes(fiberApp, authHandler)
	http.RegisterQuestionnaireRoutes(baseGroup, transaction, questionnaireHandler)

	log.Fatal(fiberApp.Listen(":3000"))
}
