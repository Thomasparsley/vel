package types

import "time"

type UpdatedAtTime struct {
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (t *UpdatedAtTime) UpdatetedNow() {
	value := time.Now()
	t.UpdatedAt = &value
}
