package vel

import "github.com/gofiber/fiber/v2"

type Route struct {
	Name     string
	Path     TypedPath
	Method   string
	Handlers []Handler
}

func registerAllRoutes(app *fiber.App, pathPrefix string, routes []Route) {
	for _, route := range routes {
		registerRoute(app, pathPrefix, route)
	}
}

func registerRoute(app *fiber.App, pathPrefix string, route Route) {
	translatedHandlers := translateAllHandlers(route.Handlers)

	app.Add(route.Method, pathPrefix+route.Path.UntypedPath(), translatedHandlers...)
}
