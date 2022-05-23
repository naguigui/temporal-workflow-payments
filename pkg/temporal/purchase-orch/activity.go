package purchaseorch

import "context"

type Repository interface {
	CreateOrder(ctx context.Context, productID string, userID string, quantity int) (orderID string, err error)
}

type Activity struct {
	Repo Repository
}

func (a *Activity) CreateOrder(ctx context.Context, productID string, userID string, quantity int) (string, error) {
	return a.Repo.CreateOrder(ctx, productID, userID, quantity)
}
