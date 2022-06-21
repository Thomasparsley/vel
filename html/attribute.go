package html

import (
	"bytes"
)

type AttributeName string

const (
	DisabledAttribute    AttributeName = "disabled"
	ForAttribute         AttributeName = "for"
	HeightAttribute      AttributeName = "height"
	IdAttribute          AttributeName = "id"
	MaxlengthAttribute   AttributeName = "maxlength"
	MinlengthAttribute   AttributeName = "minlength"
	MultipleAttribute    AttributeName = "multiple"
	NameAttribute        AttributeName = "name"
	PlaceholderAttribute AttributeName = "placeholder"
	ReadonlyAttribute    AttributeName = "readonly"
	RequiredAttribute    AttributeName = "required"
	StepAtrribute        AttributeName = "step"
	TypeAttribute        AttributeName = "type"
	ValueAttribute       AttributeName = "value"
	WidthAttribute       AttributeName = "width"
)

type Attribute struct {
	name  AttributeName
	Value string
}

func NewAttribute(name AttributeName, value ...string) Attribute {
	var v string

	if len(value) > 0 {
		v = value[0]
	}

	return Attribute{
		name:  name,
		Value: v,
	}
}

func (a Attribute) Name() AttributeName {
	return a.name
}

func (a Attribute) NameString() string {
	return string(a.name)
}

func (a Attribute) Render(buf *bytes.Buffer) {
	buf.WriteString(" ")
	buf.WriteString(a.NameString())

	if a.Value != "" {
		buf.WriteString("=\"")
		buf.WriteString(a.Value)
		buf.WriteString("\"")
	}
}
