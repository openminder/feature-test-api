package feature

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Feature struct {
	ID string
	Title string
	Description string
	Version string
	ProjectID string
	CreatedAt time.Time
	ModifiedAt time.Time
}

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
