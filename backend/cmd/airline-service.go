package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"airline-seat-booking-backend/internal/controller/booking"
	"airline-seat-booking-backend/internal/controller/health"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	formatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	}

	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.InfoLevel)

	startServer()

}

func startServer() {
	logrus.Info("Starting Hub Service Application")
	router := gin.New()
	router.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/healthz"}}),
	)

	health.InitRoutes(router)
	booking.InitRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		logrus.Error("Failed to start the Hub Service Application", err)
		return
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	logrus.Println("timeout of 5 seconds.")

	logrus.Println("Server exiting")
	os.Exit(0)
}
