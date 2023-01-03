package payment

import (
	// "github.com/benjaminli7/go-api-4iw3/product"
	// "gorm.io/gorm"
	"time"
)

type Payment struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	ProductId uint32          `json:"product_id"`
	PricePaid float64         `gorm:"type:float;not null" json:"price_paid"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

}
