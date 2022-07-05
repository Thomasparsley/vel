package types

type UintID[Self ObjectDefinition[uint64]] struct {
	Model[uint64, Self]
	ID uint64 `gorm:"primaryKey;->;index"`
}

func (id UintID[Self]) PK() uint64 {
	return id.ID
}
