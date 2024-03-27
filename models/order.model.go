package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID        uint           `gorm:"not null;" json:"user_id"`
	User          User           `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	ProductOrders []ProductOrder `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product_orders"`
	Status        string         `gorm:"not null;" json:"status"`
	PaymentID     string         `gorm:"not null;" json:"payment_id"`
	TotalPrice    uint           `gorm:"not null;" json:"total_price"`
}
