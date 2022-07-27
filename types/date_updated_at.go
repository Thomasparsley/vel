package types

import "time"

type UpdatedAtTime struct {
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (t *UpdatedAtTime) UpdatedNow() {
	value := time.Now()
	t.UpdatedAt = &value
}
