package models

import (
	"time"

	"github.com/google/uuid"
)

// gormigrateを利用している関係上、
// gormタグは実際のDBのカラムの情報をわかりやすく定義しているだけで、これをもとにDBのテーブルを作成するわけではない
type Order struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	TotalPrice   float64   `gorm:"type:numeric;not null"`
	OrderedAt    time.Time `gorm:"type:timestamp;not null"`
	Status       string    `gorm:"type:varchar(255);not null"`
	User         User
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}
