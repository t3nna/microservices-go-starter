package main

import "ride-sharing/shared/types"

// holds the types that FE sends to the BE
type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickUp"`
	Destination types.Coordinate `json:"destination"`
}
