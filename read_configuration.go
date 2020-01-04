package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func readConfiguration(c *gin.Context) {
	var config Configuration
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": objID}
	err = db.Collection.FindOne(context.TODO(), filter).Decode(&config)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.AbortWithStatus(204)
			return
		}
	}
	c.JSON(200, config)
}
