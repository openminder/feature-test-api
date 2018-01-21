package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/feature", CreateFeature)
	}

	router.Run(":8080")
}

func CreateFeature(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
