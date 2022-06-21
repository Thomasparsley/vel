package forms

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Thomasparsley/vel/validation"
)

type FormDefinition interface {
	Validate() validation.Errors
}

type Form[T FormDefinition] struct {
	Fields T
	errors validation.Errors
}

func NewForm[T FormDefinition]() Form[T] {
	var fields T

	return Form[T]{
		Fields: fields,
		errors: validation.Errors{},
	}
}

// TODO
func (f Form[F]) BindStruct() Form[F] {
	return f
}

func (f Form[F]) BindRequest(request *fiber.Ctx) Form[F] {
	request.BodyParser(&f.Fields)

	return f
}

func (f *Form[F]) IsValid() bool {
	f.errors = f.Fields.Validate()

	return len(f.errors) == 0
}
