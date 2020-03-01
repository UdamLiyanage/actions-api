package main

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createAction(c echo.Context) error {
	var action Action
	err := json.NewDecoder(c.Request().Body).Decode(&action)
	if checkError(err) {
		return c.JSON(500, err)
	}
	insertResult, err := DB.Collection.InsertOne(context.TODO(), action)
	if checkError(err) {
		return c.JSON(500, err)
	}
	action.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.JSON(201, action)
}
