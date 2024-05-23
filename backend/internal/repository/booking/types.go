package booking

import (
	"airline-seat-booking-backend/adapter/db"
	"airline-seat-booking-backend/pkg/models"
	"context"
)

type BookingRepository interface {
	GetBookedSeats(ctx context.Context, flightId int) ([]models.Seat, error)
	GetUser(ctx context.Context, userId int) (models.User, error)
}

type bookingRepository struct {
	dbClient db.DbClient
}
