package fields

import (
	"github.com/Thomasparsley/vel/database"
	"github.com/Thomasparsley/vel/hashids"
)

type IdField[Self database.ObjectDefinition[uint64]] struct {
	database.Model[uint64, Self]
	ID uint64 `gorm:"primaryKey;->;index"`
}

func (id IdField[Self]) PK() uint64 {
	return id.ID
}

func (id IdField[Self]) HashedID() (string, error) {
	return hashids.Get().Encode([]int{int(id.ID)})
}
