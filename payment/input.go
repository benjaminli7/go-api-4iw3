package payment

type InputPayment struct {
	ProductId uint32          `json:"product_id" binding:"required`
}