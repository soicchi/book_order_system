package models

import (
	"time"
)

type TimeStamps struct {
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}
