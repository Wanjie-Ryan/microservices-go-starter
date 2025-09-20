package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// domain package in clean architecture contains the models

type TripModel struct {
	ID       primitive.ObjectID
	UserID   string
	Status   string
	RideFare RideFareModel
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripsService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
}
