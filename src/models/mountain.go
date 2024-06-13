package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mountain struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"name"`
	Description      string             `json:"description" bson:"description"`
	Altitude         string             `json:"altitude" bson:"altitude"`
	HasDeathZone     bool               `json:"has_death_zone" bson:"has_death_zone"`
	Location         string             `json:"location" bson:"location"`
	FirstClimber     string             `json:"first_climber" bson:"first_climber"`
	FirstClimbedDate string             `json:"first_climbed_date" bson:"first_climbed_date"`
	MountainImg      string             `json:"mountain_img" bson:"mountain_img"`
	CountryFlagImg   string             `json:"country_flag_img" bson:"country_flag_img"`
}
