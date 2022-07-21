package types

import "time"

type CreatedAtTime struct {
	CreatedAt time.Time `gorm:"autoCreateTime;not null;index"`
}

func (t *CreatedAtTime) CreatedNow() {
	t.CreatedAt = time.Now()
}
