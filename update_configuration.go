package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateConfiguration(c *gin.Context) {
	var config Configuration
	err := json.NewDecoder(c.Request.Body).Decode(&config)
	checkError(err, c)
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(err, c)
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"configuration": config.Configuration,
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(err, c)
	c.JSON(200, res)
}
