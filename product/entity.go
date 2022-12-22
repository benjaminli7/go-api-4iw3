package product

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;unique_index" json:"name"`
	Price     float64   `gorm:"type:float;not null" json:"price"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}
