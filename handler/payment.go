package handler

import (
	"github.com/benjaminli7/go-api-4iw3/payment"
	"github.com/gin-gonic/gin"
	"net/http"
)



type paymentHandler struct {
	paymentService payment.Service
}

func NewPaymentHandler(paymentService payment.Service) *paymentHandler {
	return &paymentHandler{paymentService}
}

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

func (ph *paymentHandler) Stream(c *gin.Context) {
	// Créez un canal pour envoyer les événements de paiement à la fonction de streaming
	eventChan := make(chan payment.Payment)
	
	// Créez une goroutine pour écouter les événements de paiement et les envoyer au canal
	go func() {
		for {
			select {
			case payment := <-eventChan:
				// Envoyez un événement de paiement au client
				c.SSEvent("payment", payment)
			}
		}
	}()
	
	// Ajoutez le client au broadcaster pour recevoir les événements de paiement en temps réel
	broadcaster.Add(c.Writer)
	defer broadcaster.Remove(c.Writer)
	
	// Bloquez la goroutine jusqu'à ce que le client se déconnecte
	<-c.Writer.CloseNotify()
}

	
	
}