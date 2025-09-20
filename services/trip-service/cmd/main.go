package cmd

import (
	"context"
	"time"

	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

// import "fmt"

func main() {

	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository()
	service := service.NewService(inmemRepo)

	fare := &domain.RideFareModel{
		UserID: "34",
	}
	service.CreateTrip(ctx, fare)

	for {
		time.Sleep(time.Second)
	}

}
