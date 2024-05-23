package booking

import (
	"airline-seat-booking-backend/internal/repository/booking"
	"airline-seat-booking-backend/pkg/models"
	"context"
	"sync"
)

var service BookingService
var once sync.Once

func NewService() BookingService {
	once.Do(func() {
		service = &bookingService{
			bookingRepo: booking.NewRepository(),
		}
	})
	return service
}

func (s *bookingService) GetBookedSeats(ctx context.Context, flightId int) ([]models.Seat, error) {
	return s.bookingRepo.GetBookedSeats(ctx, flightId)
}

func (s *bookingService) GetUser(ctx context.Context, userId int) (models.User, error) {
	return s.bookingRepo.GetUser(ctx, userId)
}
