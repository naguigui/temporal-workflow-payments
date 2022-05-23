package sales

import "context"

type Repository interface {
	CreateReservation(ctx context.Context, orderID string, quantity int) error
}

type Activity struct {
	Repo Repository
}

func (a *Activity) ReserveProduct(ctx context.Context, orderID string, quantity int) error {
	return a.Repo.CreateReservation(ctx, orderID, quantity)
}
