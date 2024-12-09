package http

import (
	"strconv"

	"github.com/QBC8-Go-Group2/questionnaire/internal/media/domain"
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
	userIDInterface := c.Locals(UserIDKey)
	if userIDInterface == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not authenticated",
		})
	}

	userIDStr, ok := userIDInterface.(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	media, err := h.mediaService.Upload(c.Context(), uint(userID), file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"media_id":   media.ID,
		"media_uuid": media.UUID,
		"name":       media.Name,
		"size":       media.Size,
		"type":       media.Type,
	})
}

func (h *mediaHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid media ID",
		})
	}

	media, err := h.mediaService.GetByID(c.Context(), domain.MediaID(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Media not found",
		})
	}

	return c.JSON(fiber.Map{
		"media_id":   media.ID,
		"media_uuid": media.UUID,
		"name":       media.Name,
		"type":       media.Type,
		"size":       media.Size,
	})
}

func RegisterMediaRoutes(app *fiber.App, mediaHandler *mediaHandler) {
	api := app.Group("/api/v1")
	media := api.Group("/media")
	media.Use(JWTMiddleware())
	media.Post("/upload", mediaHandler.Upload)
	media.Get("/:id", mediaHandler.GetByID)
}
