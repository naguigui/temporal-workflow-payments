package inmemory

import (
	"context"

	"github.com/google/uuid"
)

type Repository struct {
	orders []*Order
}

type Order struct {
	ID          string
	PurchaserID string
	ProductID   string
	Quantity    int
}

func NewRepository() *Repository {
	return &Repository{
		orders: []*Order{},
	}
}

func (r *Repository) CreateOrder(ctx context.Context, productID string, userID string, quantity int) (orderID string, err error) {
	order := &Order{
		ID:          uuid.NewString(),
		PurchaserID: userID,
		ProductID:   productID,
		Quantity:    quantity,
	}
	r.orders = append(r.orders, order)
	return order.ID, nil
}
