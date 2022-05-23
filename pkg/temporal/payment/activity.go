package payment

import (
	"context"
	"fmt"
)

type Repository interface {
	CreatePaymentIntent(ctx context.Context, amount float64, currency string) (string, error)
}

type Activity struct {
	Repo Repository
}

type CreatePaymentIntentParams struct {
	Amount   float64
	Currency string
}
type CreatePaymentIntentResponse struct {
	PaymentIntentID string
}

func (a *Activity) CreatePaymentIntent(ctx context.Context, p CreatePaymentIntentParams) (*CreatePaymentIntentResponse, error) {
	id, err := a.Repo.CreatePaymentIntent(ctx, p.Amount, p.Currency)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment intent: %w", err)
	}
	return &CreatePaymentIntentResponse{
		PaymentIntentID: id,
	}, nil
}
