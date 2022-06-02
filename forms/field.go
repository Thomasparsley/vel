package forms

import (
	"bytes"
	"reflect"
	"strings"

	"github.com/Thomasparsley/vel/converter"
	"github.com/Thomasparsley/vel/html"
	"github.com/Thomasparsley/vel/validation"
)

type FieldType string

const (
	TextFieldType FieldType = "text"
)

type Field[T any] struct {
	label      string
	value      T
	atrributes []html.Attribute
}

func NewField[T any](label string, config FieldConfig[T]) Field[T] {
	return Field[T]{
		label:      label,
		atrributes: config.ToHtmlAttributes(),
	}
}

func (f Field[T]) Value() T {
	return f.value
}

func (f *Field[T]) SetValue(newValue T) {
	f.value = newValue
}

func (f *Field[T]) Attributes() []html.Attribute {
	return f.atrributes
}

func (f Field[T]) Label() string {
	return f.label
}

func (f Field[T]) LabelRenderBuffer(buf *bytes.Buffer, id ...string) {
	element := html.NewElement(html.LabelTag, f.Label(), false)

	var forID string
	for _, attribute := range f.Attributes() {
		if attribute.Name() == html.IdAttribute {
			forID = attribute.Value
		}
	}

	if len(id) > 0 {
		forID = id[0]
	}

	element.AddAttribute(html.NewAttribute(html.ForAttribute, forID))
	element.RenderIntoBuf(buf)

}

func (f Field[T]) LabelRender(id ...string) string {
	buf := new(bytes.Buffer)
	f.LabelRenderBuffer(buf, id...)

	return buf.String()
}

func (f Field[T]) FieldRenderBuffer(buf *bytes.Buffer, id ...string) {
	element := html.NewElement(html.InputTag, f.Value(), true)
	element.SetAttributes(f.Attributes())

	if len(id) > 0 {
		element.AddAttribute(html.NewAttribute(html.IdAttribute, id[0]))
	}

	element.RenderIntoBuf(buf)
}

func (f Field[T]) FieldRender(id ...string) string {
	buf := new(bytes.Buffer)
	f.FieldRenderBuffer(buf, id...)

	return buf.String()
}

func (f Field[T]) RenderBuffer(buf *bytes.Buffer, id ...string) {
	f.LabelRenderBuffer(buf, id...)
	f.FieldRenderBuffer(buf, id...)
}

func (f Field[T]) Render(id ...string) string {
	buf := new(bytes.Buffer)
	f.RenderBuffer(buf, id...)

	return buf.String()
}

type FieldConfig[T any] struct {
	id           string
	_type        FieldType
	autocomplete string
	checked      string
	disabled     bool
	max          int
	maxLength    uint
	min          int
	minLength    uint
	multiple     bool
	name         string
	pattern      string
	placeholder  string
	readonly     bool
	required     bool
	step         int
	value        T
	height       uint
	width        uint
	validators   []validation.ValidatorFunc[T]
	//accept         string
	//alt            string
	//capture        string
	//dirname        string
	//form           string
	//formAction     string
	//formEnctype    string
	//formMethod     string
	//formNoValidate string
	//formTarget     string
	//list           string
	//size           string
	//src            string
}

func (fc *FieldConfig[T]) ToHtmlAttributes() []html.Attribute {
	var attributes []html.Attribute

	if fc.id != "" {
		attributes = append(attributes, html.NewAttribute(
			html.IdAttribute,
			fc.id,
		))
	}

	if fc._type != "" {
		attributes = append(attributes, html.NewAttribute(
			html.TypeAttribute,
			string(fc._type),
		))
	}

	if fc.name != "" {
		attributes = append(attributes, html.NewAttribute(
			html.NameAttribute,
			fc.name,
		))
	}

	stringValue := converter.ToString(fc.value)
	if stringValue != "" {
		attributes = append(attributes, html.NewAttribute(
			html.ValueAttribute,
			stringValue,
		))
	}

	if fc.required {
		attributes = append(attributes, html.NewAttribute(html.RequiredAttribute))
	}

	if fc.readonly {
		attributes = append(attributes, html.NewAttribute(html.ReadonlyAttribute))
	}

	if fc.disabled {
		attributes = append(attributes, html.NewAttribute(html.DisabledAttribute))
	}

	if fc.multiple {
		attributes = append(attributes, html.NewAttribute(html.MultipleAttribute))
	}

	if fc.placeholder != "" {
		attributes = append(attributes, html.NewAttribute(
			html.PlaceholderAttribute,
			fc.placeholder,
		))
	}

	if fc.step > 0 {
		attributes = append(attributes, html.NewAttribute(
			html.StepAtrribute,
			converter.ToString(fc.step),
		))
	}

	if fc.maxLength > 0 {
		attributes = append(attributes, html.NewAttribute(
			html.MaxlengthAttribute,
			converter.ToString(fc.maxLength),
		))
	}

	if fc.minLength > 0 {
		attributes = append(attributes, html.NewAttribute(
			html.MinlengthAttribute,
			converter.ToString(fc.minLength),
		))
	}

	if fc.height > 0 {
		attributes = append(attributes, html.NewAttribute(
			html.HeightAttribute,
			converter.ToString(fc.height),
		))
	}

	if fc.width > 0 {
		attributes = append(attributes, html.NewAttribute(
			html.WidthAttribute,
			converter.ToString(fc.width),
		))
	}

	return attributes
}

// bindConfig binds all values from config to the field with reflection.
func (fc *FieldConfig[T]) bindConfig(config any) {
	configType := reflect.TypeOf(config)
	configValue := reflect.ValueOf(config)

	for i := 0; i < configType.NumField(); i++ {
		fieldName := configType.Field(i).Name
		fieldName = strings.ToLower(fieldName)

		fieldValue := configValue.FieldByName(fieldName).Interface()

		if fieldValue != nil {
			fc.bindSlot(fieldName, fieldValue)
		}
	}
}

func (fc *FieldConfig[T]) bindSlot(fieldName string, fieldValue any) {
	selfValue := reflect.ValueOf(fc)

	if selfValue.FieldByName(fieldName).IsValid() {
		selfValue.FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
	}
}
