package vel

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/schema"

	"github.com/Thomasparsley/vel/modules/user"
)

var decoder = schema.NewDecoder()

type Ctx struct {
	*fiber.Ctx
}

func NewCtx(request *fiber.Ctx) *Ctx {
	return &Ctx{
		Ctx: request,
	}
}

func (ctx Ctx) ParamsParser(out any) error {
	params := map[string][]string{}
	for k, v := range ctx.AllParams() {
		if params[k] == nil {
			params[k] = []string{}
		}

		params[k] = append(params[k], v)
	}

	return decoder.Decode(out, params)
}

func (ctx Ctx) RenderToBuffer(name string, bind any, layouts ...string) (*bytes.Buffer, error) {
	return RenderToBuffer(ctx.App(), name, bind, layouts...)
}

func (ctx Ctx) RenderToBytes(name string, bind any, layouts ...string) ([]byte, error) {
	return RenderToBytes(ctx.App(), name, bind, layouts...)
}

func (ctx Ctx) RenderToString(name string, bind any, layouts ...string) (string, error) {
	return RenderToString(ctx.App(), name, bind, layouts...)
}

func (ctx Ctx) GetLocalUser() *user.User {
	u := ctx.Locals("user")
	if u == nil {
		return nil
	}

	return u.(*user.User)
}
