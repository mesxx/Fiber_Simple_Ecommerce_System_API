package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID     uint    `gorm:"not null;" json:"user_id"`
	User       User    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	ProductID  uint    `gorm:"not null;" json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product"`
	Qty        uint    `gorm:"not null;" json:"qty"`
	TotalPrice uint    `gorm:"not null;" json:"total_price"`
}
