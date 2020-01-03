package main

import "github.com/gin-gonic/gin"

func readConfiguration(c *gin.Context) {
	c.String(200, "Read configuration function")
}
