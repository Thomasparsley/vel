package vel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet"
)

type Factory struct {
	views   fiber.Views
	modules []Module
}

// With decorator, decorate Params

func a() {

	templatingEngine := jet.New("./views", ".jet.html")

	app := fiber.New(fiber.Config{
		Views: templatingEngine,
	})

	app.Use()
}

func NewFactory() *Factory {
	return &Factory{
		views:   nil,
		modules: make([]Module, 0),
	}
}

func (factory *Factory) AddModule(module Module) {
	factory.modules = append(factory.modules, module)
}

func (factory *Factory) SetViewEngine(views fiber.Views) {
	factory.views = views
}

func (factory Factory) CreateApp() *fiber.App {
	config := fiber.Config{}

	if factory.views != nil {
		config.Views = factory.views
	}

	app := fiber.New(config)

	registerAllModules(app, factory.modules)

	return app
}
