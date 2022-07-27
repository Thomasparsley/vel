package pages

import (
	"github.com/gosimple/slug"

	"github.com/Thomasparsley/vel/database/fields"
	"github.com/Thomasparsley/vel/modules/identity"
)

const (
	TableName_Pages = "velpages"
)

type Page struct {
	fields.IdField[Page]
	Public bool   `gorm:"default:false"`
	Title  string `gorm:"size:512;index"`
	Slug   string `gorm:"size:512;index"`
	Body   string ``
	identity.CreatedByField
	identity.UpdatedByField
	fields.CreatedAtField
	fields.UpdatedAtField
}

func (Page) TableName() string {
	return TableName_Pages
}

func (p *Page) Slugify() {
	p.Slug = slug.Make(p.Title)
}

func Create(data Page) (Page, error) {
	data.Slugify()

	return Page{}.
		Object().
		Create(data)
}

func (p Page) Save() (Page, error) {
	p.Slugify()
	return Page{}.
		Object().
		Save(p)
}

func GetBySlug(slug string) (*Page, error) {
	return Page{}.
		Object().
		Where(Page{Slug: slug}).
		First()
}

func GetPublicBySlug(slug string) (*Page, error) {
	return Page{}.
		Object().
		Where(Page{Public: true, Slug: slug}).
		First()
}
