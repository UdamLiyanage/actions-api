package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/unrolled/secure"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type Database struct {
	Collection *mongo.Collection
}

var DB = Database{}

func init() {
	DB.Collection = connect()
}

func setupRouter() *gin.Engine {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:  []string{os.Getenv("ALLOWED_HOSTS")},
		SSLRedirect:   true,
		STSSeconds:    31536000,
		FrameDeny:     true,
		IsDevelopment: false,
	})

	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)

			if err != nil {
				c.Abort()
				return
			}

			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	r := gin.Default()
	auth := gin.BasicAuth(gin.Accounts{
		os.Getenv("API_AUTH_USERNAME"): os.Getenv("API_AUTH_PASSWORD"),
	})
	r.Use(auth)
	r.Use(secureFunc)

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
