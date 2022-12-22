package handler

import (
	"github.com/benjaminli7/go-api-4iw3/product"
	"github.com/gin-gonic/gin"
	"net/http"
	// "strconv"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (ph *productHandler) GetAll(c *gin.Context) {
	products, err := ph.productService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    products,
	})
}