package files

import (
	"time"
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
