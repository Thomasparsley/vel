package vel

import "github.com/gofiber/fiber/v2"

type Module struct {
	PathPrefix string
	Routes     []Route
}

func NewModule() Module {
	return Module{
		Routes: make([]Route, 0),
	}
}

// Add allows you to specify a HTTP method to register a route
func (m *Module) Add(name string, method string, path TypedPath, handlers ...Handler) Route {
	r := Route{
		Name:     name,
		Path:     path,
		Method:   method,
		Handlers: handlers,
	}

	m.Routes = append(m.Routes, r)

	return r
}

// Get registers a route for GET methods that requests a representation
// of the specified resource. Requests using GET should only retrieve data.
func (m *Module) Get(name string, path TypedPath, handlers ...Handler) Route {
	return m.Add(name, fiber.MethodGet, path, handlers...)
}

// Post registers a route for POST methods that is used to submit an entity to the
// specified resource, often causing a change in state or side effects on the server.
func (m *Module) Post(name string, path TypedPath, handlers ...Handler) Route {
	return m.Add(name, fiber.MethodPost, path, handlers...)
}

func registerAllModules(app *fiber.App, modules []Module) {
	for _, module := range modules {
		registerModule(app, module)
	}
}

func registerModule(app *fiber.App, module Module) {
	registerAllRoutes(app, module.PathPrefix, module.Routes)
}
