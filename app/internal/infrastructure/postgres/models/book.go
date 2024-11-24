package models

import (
	"time"

	"github.com/google/uuid"
)

// gormigrateを利用している関係上、
// gormタグは実際のDBのカラムの情報をわかりやすく定義しているだけで、これをもとにDBのテーブルを作成するわけではない
type Book struct {
	ID           uuid.UUID     `gorm:"type:uuid;primary_key"`
	Title        string        `gorm:"type:varchar(255);not null"`
	Author       string        `gorm:"type:varchar(255);not null"`
	Price        float64       `gorm:"type:numeric;not null"`
	Stock        int           `gorm:"type:integer;not null"`
	CreatedAt    time.Time     `gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time     `gorm:"type:timestamp;not null"`
	OrderDetails []OrderDetail `gorm:"foreignKey:BookID"`
}
