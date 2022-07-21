package types

import (
	"github.com/Thomasparsley/vel/database"
)

type Model[ID ComperableID, Self ObjectDefinition[ID]] struct{}

func (Model[ID, Self]) Object() Object[ID, Self] {
	db := database.Get()
	return NewObject[ID, Self](db)
}
