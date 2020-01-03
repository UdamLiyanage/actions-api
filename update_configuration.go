package main

import "github.com/gin-gonic/gin"

func updateConfiguration(c *gin.Context) {
	c.String(200, "Update configuration function")
}
