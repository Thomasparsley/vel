package vel

import "github.com/gofiber/fiber/v2"

type Ctx struct {
	request *fiber.Ctx
}

func NewCtx(request *fiber.Ctx) *Ctx {
	return &Ctx{
		request: request,
	}
}
