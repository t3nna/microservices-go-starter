package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string
	PackageSlug       string // ex: van, luxury, sedan
	TotalPriceInCents float64
	ExpiresAt         time.Time
}
