package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Place struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	City        Dictionary         `json:"city" bson:"city"`
	Name        Dictionary         `json:"name" bson:"name"`
	Location    Location           `json:"location" bson:"location"`
	Information Dictionary         `json:"information" bson:"information"`
	Visit       Visit              `json:"visit" bson:"visit"`
	Category    Dictionary         `json:"category" bson:"category"`
	Priority    int                `json:"priority" bson:"priority"`
}

type Dictionary map[string]string

type Location struct {
	Type        string    `validate:"required" json:"type" bson:"type"`
	Coordinates []float64 `validate:"required" json:"coordinates" bson:"coordinates"`
}

type Visit struct {
	All     float64 `json:"all,omitempty" bson:"all,omitempty"`
	Outside float64 `json:"outside,omitempty" bson:"outside,omitempty"`
}

func (p *Place) Latitude() float64 {
	return p.Location.Coordinates[1]
}

func (p *Place) Longitude() float64 {
	return p.Location.Coordinates[0]
}
