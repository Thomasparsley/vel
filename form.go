package vel

type FormDefinition interface {
	Validate() ValidationErrors
}

type Form[F FormDefinition] struct {
	fields F
	errors ValidationErrors
}

func NewForm[F FormDefinition](form F) Form[F] {
	return Form[F]{
		fields: form,
	}
}

func (f Form[F]) IsValid() bool {
	f.errors = f.fields.Validate()

	return len(f.errors) == 0
}

func (f Form[F]) Fields() F {
	return f.fields
}

func (f Form[F]) Errors() ValidationErrors {
	return f.errors
}

func (f *Form[F]) AddError(key string, message string) {
	f.errors[key] = message
}
