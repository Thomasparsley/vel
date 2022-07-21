package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func AutoUpgrade(c *fiber.Ctx) error {
	// Returns true if the client requested upgrade to the WebSocket protocol
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}

	return c.SendStatus(fiber.StatusUpgradeRequired)
}
