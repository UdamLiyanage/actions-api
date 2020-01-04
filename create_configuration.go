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
	checkError(err, c)
	insertResult, err := db.Collection.InsertOne(context.TODO(), config)
	checkError(err, c)
	config.ID = insertResult.InsertedID.(primitive.ObjectID)
	c.JSON(201, config)
}
