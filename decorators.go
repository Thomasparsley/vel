package vel

import (
	"github.com/gofiber/fiber/v2"
)

// Parse body and store to locals
func Body[B any]() func(request *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var body B

		if err := c.BodyParser(body); err != nil {
			return err
		}

		c.Locals("body", body)
		return c.Next()
	}
}

// Parse query and store to locals
func Query[Q any]() func(request *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var query Q

		if err := c.QueryParser(query); err != nil {
			return err
		}

		c.Locals("query", query)
		return c.Next()
	}
}

func Redirect(location string, status ...int) func(request *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		if err != nil {
			return err
		}

		return c.Redirect(location, status...)
	}
}
