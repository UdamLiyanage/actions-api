package main

import "github.com/gin-gonic/gin"

func deleteConfiguration(c *gin.Context) {
	c.String(200, "Delete configuration function")
}
