package vel

import "github.com/gofiber/fiber/v2"

type Handler func(ctx *Ctx) error

func translateAllHandlers(handlers []Handler) []fiber.Handler {
	result := make([]fiber.Handler, len(handlers))

	for i, handler := range handlers {
		result[i] = translateHandler(handler)
	}

	return result
}

func translateHandler(handler Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := NewCtx(c)
		return handler(ctx)
	}
}
