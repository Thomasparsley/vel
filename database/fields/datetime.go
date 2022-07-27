package fields

import "time"

type CreatedAtField struct {
	CreatedAt time.Time `gorm:"autoCreateTime;not null;index"`
}

func (t *CreatedAtField) CreatedNow() {
	t.CreatedAt = time.Now()
}

type UpdatedAtField struct {
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (t *UpdatedAtField) UpdatedNow() {
	value := time.Now()
	t.UpdatedAt = &value
}
