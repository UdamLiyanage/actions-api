package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"strconv"
)

func readAction(c echo.Context) error {
	var action Action
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	filter := bson.M{"_id": objID}
	err = collection.FindOne(context.TODO(), filter).Decode(&action)
	if checkError(err) {
		return c.JSON(404, err)
	}
	return c.JSON(200, action)
}

func readAllDeviceActions(c echo.Context) error {
	var (
		res Pager
		err error
	)
	if lim := c.QueryParam("limit"); lim == "" {
		res.Limit = 20
		if n := c.QueryParam("next"); n != "" {
			objID, err := primitive.ObjectIDFromHex(n)
			if err != nil {
				return c.JSON(500, err)
			}
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
				{"_id", bson.D{{"$gt", objID}}},
			}
			res.FirstPage = false
		} else if p := c.QueryParam("previous"); p != "" {
			objID, err := primitive.ObjectIDFromHex(p)
			if err != nil {
				return c.JSON(500, err)
			}
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
				{"_id", bson.D{{"$lt", objID}}},
			}
			res.FirstPage = false
		} else {
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
			}
			res.FirstPage = true
		}
	} else {
		res.Limit, err = strconv.ParseInt(lim, 32, 64)
		if err != nil {
			log.Println("Error Occurred: ", err)
			return c.JSON(500, err)
		}
		if n := c.QueryParam("next"); n != "" {
			objID, err := primitive.ObjectIDFromHex(c.QueryParam("next"))
			if err != nil {
				return c.JSON(500, err)
			}
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
				{"_id", bson.D{{"$gt", objID}}},
			}
			res.FirstPage = false
		} else if p := c.QueryParam("previous"); p != "" {
			objID, err := primitive.ObjectIDFromHex(p)
			if err != nil {
				return c.JSON(500, err)
			}
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
				{"_id", bson.D{{"$lt", objID}}},
			}
			res.FirstPage = false
		} else {
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
			}
			res.FirstPage = true
		}
	}
	res.URL = os.Getenv("SELF_ADDRESS") + c.Request().URL.Path
	return c.JSON(200, res.Paginate())
}
