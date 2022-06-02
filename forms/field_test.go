package forms

import (
	"testing"
)

func TestAPI(t *testing.T) {
	charField := NewCharField("Test", 10)
	charField.SetValue("1234546")

	t.Log(charField.Render())
	t.Log(charField.Render("testID"))
}
