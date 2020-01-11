package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readConfiguration(c *gin.Context) {
	var config Configuration
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(err, c)
	filter := bson.M{"_id": objID}
	err = DB.Collection.FindOne(context.TODO(), filter).Decode(&config)
	checkError(err, c)
	c.JSON(200, config)
}
