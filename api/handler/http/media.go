package http

import (
	"strconv"

	"github.com/QBC8-Go-Group2/questionnaire/internal/media/port"
	"github.com/gofiber/fiber/v2"
)

type mediaHandler struct {
	mediaService port.Service
}

func NewMediaHandler(mediaService port.Service) *mediaHandler {
	return &mediaHandler{mediaService: mediaService}
}

func (h *mediaHandler) Upload(c *fiber.Ctx) error {
	// Get user ID using the correct context key
	userIDInterface := c.Locals(UserIDKey)
	if userIDInterface == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not authenticated",
		})
	}

	// Convert the user ID to string first (as stored in middleware)
	userIDStr, ok := userIDInterface.(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	// Convert string to uint
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	// Get uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	// Upload file
	mediaID, err := h.mediaService.Upload(c.Context(), uint(userID), file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"media_id": mediaID,
	})
}

func RegisterMediaRoutes(app *fiber.App, mediaHandler *mediaHandler) {
	api := app.Group("/api/v1")
	media := api.Group("/media")
	media.Use(JWTMiddleware())
	media.Post("/upload", mediaHandler.Upload)
}
