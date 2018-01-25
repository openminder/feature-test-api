package testperiod

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TestPeriod struct {
	ID string
	From time.Time
	To time.Time
	Status string
	Tests []Test
	CreatedAt time.Time
	ModifiedAt time.Time
}

type Test struct{
	ID string
	FeatureID string
	Result bool
	CreatedAt time.Time
	ModifiedAt time.Time
}

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


