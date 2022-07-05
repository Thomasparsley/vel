package pages

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"

	"github.com/Thomasparsley/vel/types"
)

const (
	TableName_Pages = "velpages"
)

type Page struct {
	types.UintID[Page]
	Public bool   `gorm:"default:false"`
	Title  string `gorm:"size:512;index"`
	Slug   string `gorm:"size:512;index"`
	Body   string
	types.CreatedAtTime
	types.UpdatedAtTime
}

func (Page) TableName() string {
	return TableName_Pages
}

func (p *Page) Slugify() {
	p.Slug = slug.Make(p.Title)
}

func (p Page) Save(db *gorm.DB) (Page, error) {
	p.Slugify()
	return Page{}.Object(db).
		Save(p)
}

func GetBySlug(db *gorm.DB, slug string) (types.Optional[Page], error) {
	return Page{}.Object(db).
		Where(Page{Slug: slug}).
		First()
}

func GetPublicBySlug(db *gorm.DB, slug string) (types.Optional[Page], error) {
	return Page{}.Object(db).
		Where(Page{Public: true, Slug: slug}).
		First()
}

func Create(db *gorm.DB, data Page) (Page, error) {
	data.Slugify()
	return Page{}.Object(db).
		Create(data)
}
