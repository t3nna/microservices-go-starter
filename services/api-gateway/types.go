package main

// holds the types that FE sends to the BE
type previewTripRequest struct {
	UserID      string `json:"userID"`
	Pickup      string `json:"pickUp"`
	Destination string `json:"destination"`
}
