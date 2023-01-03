package main

import (
	"log"
	"os"

	_ "github.com/benjaminli7/go-api-4iw3/docs"
	"github.com/benjaminli7/go-api-4iw3/handler"
	"github.com/benjaminli7/go-api-4iw3/payment"
	"github.com/benjaminli7/go-api-4iw3/product"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title           Go API 4IW3
// @version         1.0
// @description     Go API 4IW3  is a sample application for a university project.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/go-api-4iw3?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&product.Product{}, &payment.Payment{})

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	r := gin.Default()
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.POST("/product", productHandler.Store)
	api.GET("/product", productHandler.GetAll)
	api.GET("/product/:id", productHandler.GetById)
	api.PUT("/product/:id", productHandler.Update)
	api.DELETE("/product/:id", productHandler.Delete)
	api.GET("/payment", paymentHandler.GetAll)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000")
}
