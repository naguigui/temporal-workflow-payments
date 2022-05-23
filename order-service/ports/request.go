package ports

type CreateOrderRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	UserID    string `json:"user_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type CreateOrderResponse struct {
	OrderID string `json:"order_id"`
}
