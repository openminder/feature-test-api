package feature

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddFeature(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func UpdateFeature(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func FindFeaturesByProject(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func FindFeaturesByVersions(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func FindFeatureById(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func DeleteFeature(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
