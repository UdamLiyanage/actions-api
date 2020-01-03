package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Database struct {
	Collection *mongo.Collection
}

var db = Database{Collection: connect()}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users/:id", readConfiguration)

	r.POST("/users", createConfiguration)

	r.PUT("/users/:id", updateConfiguration)

	r.DELETE("/users/:id", deleteConfiguration)
	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run())
}
