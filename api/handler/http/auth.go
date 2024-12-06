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

func (h *authHandler) Register(c *fiber.Ctx) error {
	var req domain.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	userID, err := h.authService.Register(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user_id": userID,
	})
}

func (h *authHandler) InitiateOTP(c *fiber.Ctx) error {
	var req struct {
		Email string `json:"email"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.authService.InitiateOTP(c.Context(), req.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OTP sent successfully",
	})
}

func (h *authHandler) VerifyOTP(c *fiber.Ctx) error {
	var req domain.OTPRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	token, err := h.authService.VerifyOTP(c.Context(), req)
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

	auth.Post("/sign-up", authHandler.Register)
	auth.Post("/otp/init", authHandler.InitiateOTP)
	auth.Post("/otp/verify", authHandler.VerifyOTP)
}
