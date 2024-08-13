package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID              uuid.UUID `gorm:"primary_key, unique,type:uuid, column:id,default:uuid_generate_v4()"`
	Title           string
	Price           int
	QuantityInStock int
	CreatedAt       time.Time `gorm:"type:TIMESTAMP;null;default:null"`
}

func (product *Product) TableName() string {
	return "products"
}
