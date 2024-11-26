package models

import (
	"github.com/google/uuid"
)

// gormigrateを利用している関係上、
// gormタグは実際のDBのカラムの情報をわかりやすく定義しているだけで、これをもとにDBのテーブルを作成するわけではない
type OrderDetail struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	OrderID  uuid.UUID `gorm:"type:uuid;not null"`
	BookID   uuid.UUID `gorm:"type:uuid;not null"`
	Quantity int       `gorm:"type:integer;not null"`
	Price    float64   `gorm:"type:numeric;not null"`
	Book     Book
	Order    Order
}
