package http

import (
	"github.com/QBC8-Go-Group2/questionnaire/pkg/context"
	"github.com/gofiber/fiber/v2"
	fiberLimiter "github.com/gofiber/fiber/v2/middleware/limiter"
	"gorm.io/gorm"
	"time"
)

func Limiter() fiber.Handler {
	return fiberLimiter.New(fiberLimiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        10,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	})
}

func SetUserContext(c *fiber.Ctx) error {
	c.SetUserContext(context.NewAppContext(c.UserContext()))
	return c.Next()
}

func QuestionsTransaction(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tr := db.Begin()

		context.SetDB(c.UserContext(), tr)

		err := c.Next()

		if err != nil {
			tr.Rollback()
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if err := tr.Commit().Error; err != nil {
			tr.Rollback()
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return nil
	}
}
