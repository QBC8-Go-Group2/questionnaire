package http

import (
	"github.com/gofiber/fiber/v2"
	fiberLimiter "github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func limiter() fiber.Handler {
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
