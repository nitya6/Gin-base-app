package main

import (
	"base-app/api"
	"base-app/config"
	"base-app/driver"
	"base-app/middleware"
	"base-app/utils"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @title           Portal API
// @version         1.0
// @description     This is Service Composition self service onboarding api spec.
// @contact.email  purplesea.service.composition@maersk.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath  /api/v1
// This boots up the server
func main() {

	utils.InitializeLogger()
	Logger := utils.Logger

	config.LoadAppConfig()

	metricRouter := gin.Default()
	utils.GetMetricsExporter().Use(metricRouter)

	appRouter := gin.Default()
	middleware.LoadCommonMiddleware(appRouter)
	api.LoadRoutes(appRouter)

	driver.InitializeDbConnection()

	go func() {
		Logger.Info("App metrics exported on port", zap.String("Port", config.AppConfig.Server.MetricsPort))
		if bootErr := http.ListenAndServe("0.0.0.0:"+config.AppConfig.Server.MetricsPort, metricRouter); bootErr != nil {
			log.Printf("failed to metrics server %v", bootErr)
		}
	}()

	Logger.Info("API server running on port", zap.String("Port", config.AppConfig.Server.Port))
	if bootErr := http.ListenAndServe("0.0.0.0:"+config.AppConfig.Server.Port, appRouter); bootErr != nil {
		log.Printf("failed to boot server %v", bootErr)
	}

}
