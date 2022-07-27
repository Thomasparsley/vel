package database

type Model[ID ComperableID, Self ObjectDefinition[ID]] struct{}

func (Model[ID, Self]) Object() Object[ID, Self] {
	db := Get()
	return NewObject[ID, Self](db)
}
