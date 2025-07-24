package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ride-sharing/shared/types"
)

type TripModel struct {
	ID       primitive.ObjectID // To avoid conflicts with mongo
	UserID   string
	Status   string
	RideFare *RideFareModel
}

// TODO: Learn context
type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
	GetRoute(ctx context.Context, pickup, destination *types.Coordinate) (*types.Route, error)
}
