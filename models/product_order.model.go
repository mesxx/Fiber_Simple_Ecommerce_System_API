package models

import (
	"gorm.io/gorm"
)

type ProductOrder struct {
	gorm.Model
	OrderID    uint    `gorm:"not null;" json:"order_id"`
	Order      Order   `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order"`
	ProductID  uint    `gorm:"not null;" json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product"`
	Qty        uint    `gorm:"not null;" json:"qty"`
	TotalPrice uint    `gorm:"not null;" json:"total_price"`
}
