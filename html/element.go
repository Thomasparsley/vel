package html

import (
	"bytes"

	"github.com/Thomasparsley/vel/converter"
)

type Element[V any] struct {
	tag         Tag
	value       V
	selfClosing bool
	atrributes  []Attribute
}

// NewElement creates a new element
func NewElement[V any](tag Tag, value V, selfClosing bool) Element[V] {
	return Element[V]{
		tag:         tag,
		value:       value,
		selfClosing: selfClosing,
		atrributes:  []Attribute{},
	}
}

func (e Element[V]) Tag() Tag {
	return e.tag
}

func (e Element[V]) TagName() string {
	return string(e.tag)
}

func (e *Element[V]) AddAttribute(attribute Attribute) {
	e.atrributes = append(e.atrributes, attribute)
}

func (e *Element[V]) SetAttributes(attributes []Attribute) {
	e.atrributes = attributes
}

func (e Element[V]) RenderIntoBuf(buf *bytes.Buffer) {
	value := converter.ToString(e.value)

	buf.WriteString("<")
	buf.WriteString(e.TagName())

	if e.Tag() == InputTag {
		e.AddAttribute(NewAttribute("value", value))
	}

	for _, attribute := range e.atrributes {
		attribute.Render(buf)
	}

	buf.WriteString(">")

	if !e.selfClosing {
		buf.WriteString(value)

		buf.WriteString("</")
		buf.WriteString(e.TagName())
		buf.WriteString(">")
	}
}

func (e Element[V]) Render() string {
	buf := new(bytes.Buffer)
	e.RenderIntoBuf(buf)

	return buf.String()
}
