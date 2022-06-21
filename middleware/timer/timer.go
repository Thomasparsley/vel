package timer

import (
	"fmt"
	"time"

	"github.com/Thomasparsley/vel"
)

func New() vel.Handler {
	return func(c *vel.Ctx) error {
		start := time.Now()

		err := c.Next()

		stop := time.Now()
		c.Append("Server-Timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))

		return err
	}
}
