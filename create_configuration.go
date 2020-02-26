package main

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createConfiguration(c echo.Context) error {
	var config Configuration
	err := json.NewDecoder(c.Request().Body).Decode(&config)
	if checkError(err) {
		return c.JSON(500, err)
	}
	insertResult, err := DB.Collection.InsertOne(context.TODO(), config)
	if checkError(err) {
		return c.JSON(500, err)
	}
	config.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.JSON(201, config)
}
