package main

import (
	// "io"

	"log"
	"os"

	_ "github.com/benjaminli7/go-api-4iw3/docs"
	"github.com/benjaminli7/go-api-4iw3/handler"
	"github.com/benjaminli7/go-api-4iw3/payment"
	"github.com/benjaminli7/go-api-4iw3/product"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		log.Fatal("DB", err.Error())
	}

	db.AutoMigrate(&product.Product{}, &payment.Payment{})

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	br := payment.NewBroadcaster(10)
	paymentRepository := payment.NewRepository(db, br)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService, br)

	r := gin.Default()
	api := r.Group("/api")

	api.POST("/product", productHandler.Store)
	api.GET("/product", productHandler.GetAll)
	api.GET("/product/:id", productHandler.GetById)
	api.PUT("/product/:id", productHandler.Update)
	api.DELETE("/product/:id", productHandler.Delete)

	api.POST("/payment", paymentHandler.Store)
	api.GET("/payment", paymentHandler.GetAll)
	api.GET("/payment/:id", paymentHandler.GetById)
	api.PUT("/payment/:id", paymentHandler.Update)
	api.DELETE("/payment/:id", paymentHandler.Delete)
	api.GET("/payment/stream", paymentHandler.Stream)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000")
}
