package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Configuration struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DeviceToken       string             `json:"device_token" bson:"device_token"`
	DeviceSerial      string             `json:"device_serial" bson:"device_serial"`
	ConfigurationType string             `json:"configuration_type" bson:"configuration_type"`
	Configuration     map[string]string  `json:"configuration" bson:"configuration"`
}
