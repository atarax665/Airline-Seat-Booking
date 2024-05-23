package booking

import (
	"airline-seat-booking-backend/internal/repository/booking"
	"airline-seat-booking-backend/pkg/models"
	"context"
)

type BookingService interface {
	GetBookedSeats(ctx context.Context, flightId int) ([]models.Seat, error)
	GetUser(ctx context.Context, userId int) (models.User, error)
}

type bookingService struct {
	bookingRepo booking.BookingRepository
}
