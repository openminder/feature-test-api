package testperiod

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTestPeriod(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func UpdateTestPeriod(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func FindTestPeriodByProject(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func AddFeaturesToTestPeriod(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func UpdateTestResultOfTestPeriod(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func DeleteFeatureFromTestPeriod(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func FindTestPeriodById(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func DeleteTestPeriod(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}


