package main

import "ride-sharing/shared/types"

type previewTripRequest struct {
	UserID  string           `json:"userID"`
	Pickup  types.Coordinate `json:"pickup"`
	Dropoff types.Coordinate `json:"dropoff"`
}
