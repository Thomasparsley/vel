package types

import "gorm.io/gorm"

type Model[ID ComperableID, Self ObjectDefinition[ID]] struct{}

func (Model[ID, Self]) Object(db *gorm.DB) Object[ID, Self] {
	return NewObject[ID, Self](db)
}
