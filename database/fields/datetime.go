package fields

import "time"

type CreatedAtField struct {
	CreatedAt time.Time `gorm:"autoCreateTime;not null;index"`
}

func (t *CreatedAtField) CreatedNow() {
	t.CreatedAt = time.Now()
}

type UpdatedAtField struct {
	UpdatedAt *time.Time `gorm:"autoUpdateTime;index"`
}

func (t *UpdatedAtField) UpdatedNow() {
	value := time.Now()
	t.UpdatedAt = &value
}
