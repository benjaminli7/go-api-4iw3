package handler

import (
	"io"
	"net/http"
	"strconv"

	"github.com/benjaminli7/go-api-4iw3/payment"
	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentService payment.Service
	br             payment.Broadcaster
}

func NewPaymentHandler(paymentService payment.Service, br payment.Broadcaster) *paymentHandler {
	return &paymentHandler{paymentService, br}
}

// Stream godoc
// @Summary      Stream
// @Description  Stream
// @Tags         stream
// @Accept       json
// @Produce      json
// @Param        q    body     array  true  "stream"
// @Success      200  {array}   bool
// @Failure      400  {object}  Response
// @Router       /payment/stream [get]
func (ph *paymentHandler) Stream(c *gin.Context) {
	listener := make(chan interface{})
	ph.br.Register(listener)
	defer ph.br.Unregister(listener)

	clientGone := c.Request.Context().Done()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case message := <-listener:
			serviceMsg, ok := message.(payment.Payment)
			// fmt.Println(message)
			if !ok {
				c.SSEvent("message", message)
				return false
			}
			// fmt.Println(message)
			c.SSEvent("message", serviceMsg)
			return true
		}
	})
}

// PostPayment godoc
// @Summary      Post a new payment
// @Description  add payment
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id   body     	int  true  "Add payment"
// @Success      200  {object}  payment.InputPayment
// @Failure      400  {object}  Response
// @Router       /payment [post]
func (ph *paymentHandler) Store(c *gin.Context) {
	var input payment.InputPayment

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

	newPayment, err := ph.paymentService.Store(input)
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
		Message: "New payment created",
		Data:    newPayment,
	}
	c.JSON(http.StatusCreated, response)
}

// ListPayments godoc
// @Summary      List payments
// @Description  get all payments
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "get all payments"
// @Success      200  {array}   payment.Payment
// @Failure      400  {object}  Response
// @Router       /payment [get]
func (ph *paymentHandler) GetAll(c *gin.Context) {
	payments, err := ph.paymentService.GetAll()
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
		Data:    payments,
	})
}

// ShowPayment godoc
// @Summary      Show a payment
// @Description  get payment by ID
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Payment ID"
// @Success      200  {object}  payment.Payment
// @Failure      400  {object}  Response
// @Router       /payment/{id} [get]
func (ph *paymentHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	payment, err := ph.paymentService.GetById(id)
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
		Data:    payment,
	})
}

// UpdatePayment godoc
// @Summary      Update a payment
// @Description  update payment by ID
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id   body     	int  true  "update payments"
// @Success      200  {object}  payment.InputPayment
// @Failure      400  {object}  Response
// @Router       /payment/{id} [PUT]
func (ph *paymentHandler) Update(c *gin.Context) {
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
	var input payment.InputPayment
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

	payment, err := ph.paymentService.Update(id, input)
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
		Message: "Payment successfully updated",
		Data:    payment,
	}
	c.JSON(http.StatusCreated, response)
}

// DeletePayment godoc
// @Summary      Delete a payment
// @Description  delete payment by ID
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Payment ID"
// @Success      200  {object} 	string
// @Failure      400  {object}  Response
// @Router       /payment/{id} [DELETE]
func (ph *paymentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	err = ph.paymentService.Delete(id)
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
		Message: "Payment successfully deleted",
	})
}

// func (ph *paymentHandler) Stream(c *gin.Context) {
// 	// Cr??ez un canal pour envoyer les ??v??nements de paiement ?? la fonction de streaming
// 	eventChan := make(chan payment.Payment)

// 	// Cr??ez une goroutine pour ??couter les ??v??nements de paiement et les envoyer au canal
// 	go func() {
// 		for {
// 			select {
// 			case payment := <-eventChan:
// 				// Envoyez un ??v??nement de paiement au client
// 				c.SSEvent("payment", payment)
// 			}
// 		}
// 	}()

// 	// Ajoutez le client au broadcaster pour recevoir les ??v??nements de paiement en temps r??el
// 	broadcaster.Add(c.Writer)
// 	defer broadcaster.Remove(c.Writer)

// 	// Bloquez la goroutine jusqu'?? ce que le client se d??connecte
// 	<-c.Writer.CloseNotify()
// }
