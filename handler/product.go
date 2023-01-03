package handler

import (
	"net/http"

	"strconv"

	"github.com/benjaminli7/go-api-4iw3/product"
	"github.com/gin-gonic/gin"
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

// PostProduct godoc
// @Summary      Post a new product
// @Description  add product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   body     	int  true  "Add product"
// @Success      200  {object}  product.InputProduct
// @Failure      400  {object}  Response
// @Router       /product [post]
func (ph *productHandler) Store(c *gin.Context) {
	var input product.InputProduct

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newProduct, err := ph.productService.Store(input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &Response{
		Success: true,
		Message: "New product created",
		Data:    newProduct,
	}
	c.JSON(http.StatusCreated, response)
}

// ListProducts godoc
// @Summary      List products
// @Description  get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "get all products"
// @Success      200  {array}   product.Product
// @Failure      400  {object}  Response
// @Router       /product [get]
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

// ShowProduct godoc
// @Summary      Show an product
// @Description  get product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  product.Product
// @Failure      400  {object}  Response
// @Router       /product/{id} [get]
func (ph *productHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	product, err := ph.productService.GetById(id)
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
		Data:    product,
	})
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  update product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   body     	int  true  "update product"
// @Success      200  {object}  product.InputProduct
// @Failure      400  {object}  Response
// @Router       /product/{id} [PUT]
func (ph *productHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	// Get json body
	var input product.InputProduct
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := ph.productService.Update(id, input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &Response{
		Success: true,
		Message: "Product successfully updated",
		Data:    product,
	}
	c.JSON(http.StatusCreated, response)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  delete product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object} 	string
// @Failure      400  {object}  Response
// @Router       /product/{id} [DELETE]
func (ph *productHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	err = ph.productService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Product successfully deleted",
	})
}
