package models

import (
	"time"

	"github.com/google/uuid"
)

// gormigrateを利用している関係上、
// gormタグは実際のDBのカラムの情報をわかりやすく定義しているだけで、これをもとにDBのテーブルを作成するわけではない
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Username  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null;unique"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
	Orders    []Order   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
