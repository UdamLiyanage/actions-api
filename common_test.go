package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	DB.Collection = connect()
	// Run the other tests
	os.Exit(m.Run())
}

func newRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/configurations/:id", readConfiguration)
	r.POST("/configurations", createConfiguration)
	r.DELETE("/configurations/:id", deleteConfiguration)
	r.PUT("/configurations/:id", updateConfiguration)
	return r
}
