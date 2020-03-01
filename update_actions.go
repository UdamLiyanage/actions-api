package main

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateAction(c echo.Context) error {
	var action Action
	err := json.NewDecoder(c.Request().Body).Decode(&action)
	if checkError(err) {
		return c.JSON(500, err)
	}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"device_token":  action.DeviceToken,
			"device_serial": action.DeviceSerial,
			"action_type":   action.ActionType,
			"configuration": action.ActionConfig,
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	if checkError(err) {
		return c.JSON(500, err)
	}
	return c.JSON(200, res)
}
