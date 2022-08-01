package files

import (
	"github.com/Thomasparsley/vel/database/fields"
	"github.com/Thomasparsley/vel/modules/identity"
)

const (
	TableName_Files = "vel__files"
)

type File struct {
	fields.IdField[File]
	Public      bool   `gorm:"default:false"`
	Filename    string `gorm:"size:2048"`
	Size        uint64 ``
	ContentType string `gorm:"size:256"`
	identity.UploadedByField
	identity.UpdatedByField
	fields.CreatedAtField
	fields.UpdatedAtField
}

func (File) TableName() string {
	return TableName_Files
}
