package pages

import (
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"

	"github.com/Thomasparsley/vel/database"
	"github.com/Thomasparsley/vel/structs/optional"
	"github.com/Thomasparsley/vel/structs/result"
)

const (
	TableName_Pages = "velpages"
)

type Page struct {
	ID        uint64                       `gorm:"primaryKey;->;index"`
	Public    bool                         `gorm:"default:false"`
	Title     string                       `gorm:"size:512;index"`
	Slug      string                       `gorm:"size:512;index"`
	Body      string                       ``
	CreatedAt time.Time                    `gorm:"autoCreateTime;not null;index"`
	UpdatedAt optional.Optional[time.Time] `gorm:"autoUpdateTime"`
}

func (p Page) PK() uint64 {
	return p.ID
}

func (Page) TableName() string {
	return TableName_Pages
}

func (Page) Object(db *gorm.DB) database.Object[uint64, Page] {
	return database.NewObject[uint64, Page](db)
}

func (p *Page) slugify() {
	p.Slug = slug.Make(p.Title)
}

func (p Page) Save(db *gorm.DB) result.Result[Page] {
	p.slugify()
	return Page{}.Object(db).Save(p)
}

func GetBySlug(db *gorm.DB, slug string) result.Result[optional.Optional[Page]] {
	return Page{}.Object(db).
		Where(Page{Slug: slug}).
		First()
}

func GetPublicBySlug(db *gorm.DB, slug string) result.Result[optional.Optional[Page]] {
	return Page{}.Object(db).
		Where(Page{Public: true, Slug: slug}).
		First()
}

func Create(db *gorm.DB, data Page) result.Result[Page] {
	data.slugify()
	return Page{}.Object(db).Create(data)
}
