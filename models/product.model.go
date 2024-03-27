package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductOrders []ProductOrder `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product_orders"`
	Carts         []Cart         `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"carts"`
	Title         string         `gorm:"not null;" json:"title"`
	Qty           uint           `gorm:"not null;" json:"qty"`
	Price         uint           `gorm:"not null;" json:"price"`
	Description   sql.NullString `gorm:"default:NULL;" json:"description"`
	Image         sql.NullString `gorm:"default:NULL;" json:"image"`
}
