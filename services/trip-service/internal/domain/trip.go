package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID     primitive.ObjectID // To avoid conflicts with mongo
	UserID string
	Status string
	RideFareModel
}

// TODO: Lern context
type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
}
