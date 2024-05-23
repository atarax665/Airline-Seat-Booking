package booking

import (
	"airline-seat-booking-backend/internal/service/booking"

	"github.com/gin-gonic/gin"
)

type BookingController interface {
	GetBookedSeats(c *gin.Context)
	GetUser(c *gin.Context)
}

type bookingController struct {
	bookingService booking.BookingService
}
