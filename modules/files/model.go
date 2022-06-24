package files

import (
	"time"

	"gorm.io/gorm"

	"github.com/Thomasparsley/vel/database"
)

const (
	TableName_Files = "velfiles"
)

type File struct {
	ID          string     `gorm:"primaryKey;->;index;type:uuid"`
	Public      bool       `gorm:"default:false"`
	Filename    string     `gorm:"size:2048"`
	Size        uint64     ``
	ContentType string     `gorm:"size:256"`
	CreatedAt   time.Time  `gorm:"autoCreateTime;not null;index"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime"`
}

func (f File) PK() string {
	return f.ID
}

func (File) TableName() string {
	return TableName_Files
}

func (File) Object(db *gorm.DB) database.Object[string, File] {
	return database.NewObject[string, File](db)
}
