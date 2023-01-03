package main

import (
	// "io"
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

	api.POST("/payment", paymentHandler.Store)
	api.GET("/payment", paymentHandler.GetAll)
	api.GET("/payment/:id", paymentHandler.GetById)
	api.PUT("/payment/:id", paymentHandler.Update)
	api.DELETE("/payment/:id", paymentHandler.Delete)

	// api.GET("/payment/stream", func(c *gin.Context) {
	// 	ch := broadcaster.Subscribe()
	// 	defer close(ch)

	// 	c.Stream(func(w io.Writer) bool {
	// 		payment, ok := <-ch
	// 		if !ok {
	// 			return false
	// 		}
	// 		c.SSEvent("payment", payment)
	// 		return true
	// 	})
	// })
	// go func() {
	// 	clients := make(map[chan Payment]struct{})
	// 	for {
	// 		select {
	// 		case ch := <-b.subscribe:
	// 			clients[ch] = struct{}{}
	// 		case ch := <-b.unsubscribe:
	// 			delete(clients, ch)
	// 			close(ch)
	// 		case payment := <-b.payments:
	// 			for ch := range clients {
	// 				ch <- payment
	// 			}
	// 		}
	// 	}
	// }()

	r.Run(":3000")
}
