package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `gorm:"unique;not null;" json:"name"`
	Email    string  `gorm:"unique;not null;" json:"email"`
	Password string  `gorm:"not null;" json:"password"`
	Carts    []Cart  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"carts"`
	Orders   []Order `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`
}
