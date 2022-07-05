package types

import (
	"time"
)

type CreatedAtTime struct {
	CreatedAt time.Time `gorm:"autoCreateTime;not null;index"`
}

func (t *CreatedAtTime) CreatedNow() {
	t.CreatedAt = time.Now()
}

type UpdatedAtTime struct {
	UpdatedAt Optional[time.Time] `gorm:"autoUpdateTime"`
}

func (t *UpdatedAtTime) UpdatetedNow() {
	t.UpdatedAt.state = OptionalState_Set
	t.UpdatedAt.value = time.Now()
}
