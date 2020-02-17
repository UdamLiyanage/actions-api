package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Database struct {
	Collection *mongo.Collection
}

var DB = Database{}

func init() {
	DB.Collection = connect()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/configurations/:id", readConfiguration)

	r.POST("/configurations/create", createConfiguration)

	r.PUT("/configurations/:id", updateConfiguration)

	r.DELETE("/configurations/:id", deleteConfiguration)
	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run(":8003"))
}
