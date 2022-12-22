package payment

import (
	"time"
	"github.com/benjaminli7/go-api-4iw3/product"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID        uint32          `gorm:"primary_key;auto_increment" json:"id"`
	ProductId uint32          `json:"product_id"`
	Product   product.Product `json:"product"`
	PricePaid float64         `gorm:"type:float;not null" json:"price_paid"`
	CreatedAt time.Time       `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}
