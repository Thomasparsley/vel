package types

type UiidID[Self ObjectDefinition[string]] struct {
	Model[string, Self]
	ID string `gorm:"primaryKey;->;index;type:uuid"`
}

func (id UiidID[Self]) PK() string {
	return id.ID
}
