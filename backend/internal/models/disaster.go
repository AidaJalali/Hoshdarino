package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DisasterType string

const (
	Earthquake DisasterType = "earthquake"
	Flood      DisasterType = "flood"
	Fire       DisasterType = "fire"
	Storm      DisasterType = "storm"
	Snow       DisasterType = "snow"
)

type Disaster struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type        DisasterType       `bson:"type" json:"type"`
	Location    Location           `bson:"location" json:"location"`
	Magnitude   float64            `bson:"magnitude,omitempty" json:"magnitude,omitempty"`
	Description string             `bson:"description" json:"description"`
	RiskLevel   int                `bson:"riskLevel" json:"riskLevel"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type Location struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}
