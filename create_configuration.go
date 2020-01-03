package main

import "github.com/gin-gonic/gin"

func createConfiguration(c *gin.Context) {
	c.String(200, "Create configuration function")
}
