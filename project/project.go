package project

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Project struct {
	ID string
	Title string
	Description string
	CreatedAt time.Time
	ModifiedAt time.Time
}

func AddProject(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func UpdateProject(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func FindAllProjects(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func FindProjectById(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func DeleteProject(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
