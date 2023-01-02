package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(255);not null;unique_index" json:"name"`
	Price     float64   `gorm:"type:float;not null" json:"price"`
}
