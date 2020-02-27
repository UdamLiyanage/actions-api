package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readConfiguration(c echo.Context) error {
	var config Configuration
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	filter := bson.M{"_id": objID}
	err = DB.Collection.FindOne(context.TODO(), filter).Decode(&config)
	if checkError(err) {
		return c.JSON(404, err)
	}
	return c.JSON(200, config)
}

func readAllConfigurations(c echo.Context) error {
	var configs []Configuration
	deviceID := c.Param("id")
	filter := bson.D{{"device_token", deviceID}}
	cur, err := DB.Collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	for cur.Next(context.TODO()) {
		var config Configuration
		err := cur.Decode(&config)
		if err != nil {
			panic(err)
		}
		configs = append(configs, config)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	return c.JSON(200, configs)
}
