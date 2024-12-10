package http

import (
	"log"
	"path/filepath"
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
	log.Printf("User ID from context: %v", c.Locals(UserIDKey))

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

func (h *mediaHandler) Download(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing media UUID",
		})
	}

	// Still get the userID from JWT to ensure user is logged in
	userIDStr := c.Locals(UserIDKey).(string)
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	media, err := h.mediaService.VerifyFileAccess(c.Context(), domain.MediaUUID(uuid), uint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Media not found",
		})
	}

	return c.Download(string(media.Path), filepath.Base(string(media.Path)))
}

func RegisterMediaRoutes(api fiber.Router, mediaHandler *mediaHandler) {
	media := api.Group("/media")
	media.Use(JWTMiddleware())
	media.Post("/upload", mediaHandler.Upload)
	media.Get("/download/:uuid", mediaHandler.Download)
}
