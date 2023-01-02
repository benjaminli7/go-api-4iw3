package payment

import (
	"github.com/benjaminli7/go-api-4iw3/product"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ProductId uint32          `json:"product_id"`
	Product   product.Product `json:"product"`
	PricePaid float64         `gorm:"type:float;not null" json:"price_paid"`
}
