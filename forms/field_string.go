package forms

type StringField struct {
	Field[string]
}

func NewStringField(value string) StringField {
	config := FieldConfig[string]{}

	return StringField{
		Field: NewField(value, config),
	}
}
