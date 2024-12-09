package http

import (
	"github.com/QBC8-Go-Group2/questionnaire/internal/auth/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/auth/port"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService port.Service
}

func NewAuthHandler(authService port.Service) *authHandler {
	return &authHandler{authService: authService}
}

func (h *authHandler) InitiateRegister(c *fiber.Ctx) error {
	var req domain.InitiateRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.authService.InitiateRegister(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OTP sent successfully",
	})
}

func (h *authHandler) CompleteRegister(c *fiber.Ctx) error {
	var req domain.CompleteRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.authService.CompleteRegister(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registration completed successfully",
	})
}

func (h *authHandler) InitiateLogin(c *fiber.Ctx) error {
	var req domain.InitiateLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.authService.InitiateLogin(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OTP sent successfully",
	})
}

func (h *authHandler) CompleteLogin(c *fiber.Ctx) error {
	var req domain.CompleteLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	token, err := h.authService.CompleteLogin(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func RegisterAuthRoutes(app *fiber.App, authHandler *authHandler) {
	api := app.Group("/api/v1")
	auth := api.Group("/auth")

	// Registration endpoints
	auth.Post("/register/init", authHandler.InitiateRegister)
	auth.Post("/register/complete", authHandler.CompleteRegister)

	// Login endpoints
	auth.Post("/login/init", authHandler.InitiateLogin)
	auth.Post("/login/complete", authHandler.CompleteLogin)

	// Protected route example
	protected := api.Group("/protected")
	protected.Use(JWTMiddleware())
	protected.Get("/profile", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id": c.Locals(UserIDKey),
			"email":   c.Locals(EmailKey),
			"role":    c.Locals(RoleKey),
		})
	})

}
