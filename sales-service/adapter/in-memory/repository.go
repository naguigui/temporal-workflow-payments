package inmemory

import "context"

type Repository struct {
	reservations []*Reservation
}

func NewRepository() *Repository {
	return &Repository{
		reservations: []*Reservation{},
	}
}

type Reservation struct {
	OrderID  string
	Quantity int
}

func (r *Repository) CreateReservation(ctx context.Context, orderID string, quantity int) error {
	r.reservations = append(r.reservations, &Reservation{
		OrderID:  orderID,
		Quantity: quantity,
	})
	return nil
}
