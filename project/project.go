package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
