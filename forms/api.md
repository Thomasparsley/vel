
https://docs.djangoproject.com/en/4.0/topics/forms/
https://wtforms.readthedocs.io/en/3.0.x/

```go

import "github.com/Thomasparsley/vel/forms"

type ExampleForm struct {
	FirstName forms.StringField   `form:"first_name"`
	LastName  forms.StringField   `form:"last_name"`
	Age       forms.IntField      `form:"age"`
	Email     forms.EmailField    `form:"email"`
	Password  forms.PasswordField `form:"password"`
	Sex       forms.CharField     `form:"sex"`
	Slug      forms.SlugField     `form:"slug`
	IsActive  forms.BooleanField  `form:"is_active"`
	About     forms.TextField     `form:"about"`
	Time      forms.TimeField     `form:"time"`
	Date      forms.DateField     `form:"date"`
	DateTime  forms.DateTimeField `form:"date_time"`
	File      forms.FileField     `form:"file"`
	Image     forms.ImageField    `form:"image"`
	IPV4      forms.IPV4Field     `form:"ipv4"`
	IPV6      forms.IPV6Field     `form:"ipv6"`
	Files     forms.FilesField    `form:"files"`
}

func NewExampleForm(rawForm) ExampleForm {
    form := ExampleForm{
		FirstName: forms.NewStringField("First Name", forms.StringFieldConfig{}),
		LastName:  forms.NewStringField("Last Name", forms.StringFieldConfig{}),
		Age:       forms.NewIntField("Age", forms.IntFieldConfig{}),
		Email:     forms.NewEmailField("Email", forms.EmailFieldConfig{}),
		Password:  forms.NewPasswordField("Password", forms.PasswordFieldConfig{}),
		...
	}

    forms.BindValues(form, values)

	return form
}
```
