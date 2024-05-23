package booking

import (
	"airline-seat-booking-backend/internal/service/booking"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var controller BookingController
var once sync.Once

func NewController() BookingController {
	once.Do(func() {
		controller = &bookingController{
			bookingService: booking.NewService(),
		}
	})
	return controller
}

func InitRoutes(router *gin.Engine) {
	controller := NewController()

	bookingRoute := router.
		Group("/v1")
	{
		bookingRoute.GET("/get", controller.GetBookedSeats)
		bookingRoute.GET("/user", controller.GetUser)
	}
}

func (c *bookingController) GetBookedSeats(ctx *gin.Context) {
	flightId := ctx.Query("flightId")
	flightIdInt, err := strconv.Atoi(flightId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "flightId must be an integer"})
		return
	}
	bookedSeats, err := c.bookingService.GetBookedSeats(ctx, flightIdInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"bookedSeats": bookedSeats})
}

func (c *bookingController) GetUser(ctx *gin.Context) {
	userId := ctx.Query("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId must be an integer"})
		return
	}
	user, err := c.bookingService.GetUser(ctx, userIdInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"name": user.Name})
}
