package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createConfiguration(c *gin.Context) {
	var config Configuration
	err := json.NewDecoder(c.Request.Body).Decode(&config)
	if err != nil {
		panic(err)
	}
	insertResult, err := db.Collection.InsertOne(context.TODO(), config)
	if err != nil {
		panic(err)
	}
	config.ID = insertResult.InsertedID.(primitive.ObjectID)
	c.JSON(201, config)
}
