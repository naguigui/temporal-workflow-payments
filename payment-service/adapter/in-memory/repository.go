package inmemory

import (
	"context"

	"github.com/google/uuid"
)

type Repository struct {
	paymentIntents []*PaymentIntent
}

func NewRepository() *Repository {
	return &Repository{
		paymentIntents: []*PaymentIntent{},
	}
}

type PaymentIntent struct {
	ID            string
	PurchasePrice float64
	Currency      string
}

func (r *Repository) CreatePaymentIntent(ctx context.Context, purchasePrice float64, currency string) (string, error) {
	paymentIntent := &PaymentIntent{
		ID:            uuid.NewString(),
		PurchasePrice: purchasePrice,
		Currency:      currency,
	}
	r.paymentIntents = append(r.paymentIntents, paymentIntent)
	return paymentIntent.ID, nil
}
