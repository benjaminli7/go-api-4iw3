package main

import (
	"log"
	"os"

	"github.com/benjaminli7/go-api-4iw3/handler"
	"github.com/benjaminli7/go-api-4iw3/payment"
	"github.com/benjaminli7/go-api-4iw3/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/go-api-4iw3?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&payment.Payment{}, &product.Product{})

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	broadcaster := payment.NewBroadcaster()
	paymentRepository := payment.NewRepository(db, broadcaster)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	r := gin.Default()
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.GET("/product", productHandler.GetAll)
	api.GET("/payment", paymentHandler.GetAll)

	r.Run(":3000")
}
