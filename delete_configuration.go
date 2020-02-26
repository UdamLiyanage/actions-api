package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func deleteConfiguration(c echo.Context) error {
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	filter := bson.D{
		bson.E{Key: "_id", Value: objID},
	}
	_, err = DB.Collection.DeleteOne(context.TODO(), filter, opts)
	if checkError(err) {
		return c.JSON(500, err)
	}
	return c.JSON(404, nil)
}
