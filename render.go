package vel

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
)

func RenderToBuffer(app *fiber.App, name string, bind any, layouts ...string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	if err := app.Config().Views.Render(buf, name, bind, layouts...); err != nil {
		return nil, err
	}

	return buf, nil
}

func RenderToBytes(app *fiber.App, name string, bind any, layouts ...string) ([]byte, error) {
	buf, err := RenderToBuffer(app, name, bind, layouts...)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func RenderToString(app *fiber.App, name string, bind any, layouts ...string) (string, error) {
	buf, err := RenderToBuffer(app, name, bind, layouts...)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
