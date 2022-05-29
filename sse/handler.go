package sse

import (
	"bufio"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func Handler(handler func(c *fiber.Ctx, w *bufio.Writer)) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		c.
			Context().
			SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
				handler(c, w)
			}))

		return nil
	}
}
