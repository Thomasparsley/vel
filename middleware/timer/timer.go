package timer

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		stop := time.Now()
		c.Append("Server-Timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))

		return err
	}
}
