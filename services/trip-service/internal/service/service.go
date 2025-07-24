package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

type service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}
func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserID:   fare.UserID,
		Status:   "pending",
		RideFare: fare,
	}
	return s.repo.CreateTrip(ctx, t)
}

func (s *service) GetRoute(ctx context.Context, pickup, destination *types.Coordinate) (*types.Route, error) {
	return &types.Route{
		Distance: 0,
		Duration: 0,
		Geometry: nil,
	}, nil
}
