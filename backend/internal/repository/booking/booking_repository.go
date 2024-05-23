package booking

import (
	"airline-seat-booking-backend/adapter/db"
	"airline-seat-booking-backend/pkg/models"
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

var bookingRepo BookingRepository
var once sync.Once

func NewRepository() BookingRepository {
	once.Do(func() {
		bookingRepo = &bookingRepository{
			dbClient: db.GetClient(),
		}
	})
	return bookingRepo
}

func (r *bookingRepository) GetBookedSeats(ctx context.Context, flightId int) ([]models.Seat, error) {
	var bookedSeats []models.Seat
	err := r.dbClient.GetDb().Where("trip_id = ? AND user_id IS NOT NULL", flightId).Find(&bookedSeats).Error
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get booked seats for flight %d", flightId))
	}
	fmt.Println(bookedSeats)
	return bookedSeats, nil
}

func (r *bookingRepository) GetUser(ctx context.Context, userId int) (models.User, error) {
	var user models.User
	err := r.dbClient.GetDb().Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return models.User{}, errors.Wrap(err, fmt.Sprintf("failed to get user with id %d", userId))
	}
	return user, nil
}
