package files

import (
	"github.com/Thomasparsley/vel/modules/identity"
	"github.com/Thomasparsley/vel/types"
)

const (
	TableName_Files = "velfiles"
)

type File struct {
	types.IdField[File]
	Public       bool          `gorm:"default:false"`
	Filename     string        `gorm:"size:2048"`
	Size         uint64        ``
	ContentType  string        `gorm:"size:256"`
	UploadedByID uint64        ``
	UploadedBy   identity.User ``
	types.CreatedAtTime
	types.UpdatedAtTime
}

func (File) TableName() string {
	return TableName_Files
}
