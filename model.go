package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Paginator interface {
		Paginate() PaginateResult
	}

	Action struct {
		ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		DeviceToken  string             `json:"device_token" bson:"device_token"`
		DeviceSerial string             `json:"device_serial" bson:"device_serial"`
		ActionType   string             `json:"action_type" bson:"action_type"`
		ActionConfig map[string]string  `json:"action" bson:"action"`
	}

	Pager struct {
		Limit     int64
		URL       string
		FirstPage bool
		Filter    bson.D
	}

	PaginateResult struct {
		NextPage string   `json:"next_page"`
		PrevPage string   `json:"previous_page"`
		Data     []Action `json:"data"`
	}
)
