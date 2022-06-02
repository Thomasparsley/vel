package forms

import "github.com/Thomasparsley/vel/validation"

type CharField struct {
	Field[string]
}

type CharFieldConfig struct {
	ID         string
	Required   bool
	MinLength  uint
	Readonly   bool
	Validators []validation.ValidatorFunc[string]
}

func NewCharField(label string, maxLenght uint, config ...CharFieldConfig) CharField {
	fieldConfig := FieldConfig[string]{
		_type:     TextFieldType,
		maxLength: maxLenght,
	}

	if len(config) > 0 {
		fieldConfig.bindConfig(config)
	}

	return CharField{
		Field: NewField(label, fieldConfig),
	}
}
