package main

import (
	"context"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
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

	mux := http.NewServeMux()

	httphandler := h.HttpHandler{Service: svc}

	mux.HandleFunc("POST /preview", httphandler.HandleTripPreview)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Http server error: %v", err)
	}

}
