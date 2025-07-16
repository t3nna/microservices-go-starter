package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	fare := &domain.RideFareModel{
		UserID: "42",
	}
	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(t)

	// keeps the program running
	for {
		time.Sleep(time.Second)
	}

}
