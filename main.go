package main

import (
	"github.com/gin-gonic/gin"
	"github.com/openminder/feature-test-api/feature"
	"github.com/openminder/feature-test-api/project"
	"github.com/openminder/feature-test-api/testperiod"
)

func main() {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Simple group: v1
	v1 := r.Group("/v1")
	v1.Use()
	{
		featureGroup := v1.Group("/feature")
		featureGroup.POST("/", feature.AddFeature)
		featureGroup.PUT("/", feature.UpdateFeature)
		featureGroup.GET("/findByProject", feature.FindFeaturesByProject)
		featureGroup.GET("/findByVersions", feature.FindFeaturesByVersions)
		featureGroup.GET("/findById/:featureID", feature.FindFeatureById)
		featureGroup.DELETE("/:featureID", feature.DeleteFeature)

		projectGroup := v1.Group("/project")
		projectGroup.POST("/", project.AddProject)
		projectGroup.PUT("/", project.UpdateProject)
		projectGroup.GET("/", project.FindAllProjects)
		projectGroup.GET("/findById/:featureID", project.FindProjectById)
		projectGroup.DELETE("/:featureID", project.DeleteProject)

		testPeriodGroup := v1.Group("/test-period")
		testPeriodGroup.POST("/", testperiod.AddTestPeriod)
		testPeriodGroup.PUT("/", testperiod.UpdateTestPeriod)
		testPeriodGroup.GET("/findByProject", testperiod.FindTestPeriodByProject)
		testPeriodGroup.POST("/addFeatures", testperiod.AddFeaturesToTestPeriod)
		testPeriodGroup.PATCH("/updateResult/:testPeriodID/:testID", testperiod.UpdateTestResultOfTestPeriod)
		testPeriodGroup.PATCH("/removeFeatures", testperiod.DeleteFeatureFromTestPeriod)
		testPeriodGroup.GET("/findById/:testPeriodId", testperiod.FindTestPeriodById)
		testPeriodGroup.DELETE("/:testPeriodId", testperiod.DeleteTestPeriod)

	}
	r.Run(":8080")
}
