package main

import (
	"base-app/api"
	"base-app/config"
	"base-app/driver"
	"base-app/middleware"
	"base-app/utils"
	"base-app/services"
	"base-app/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
    "context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
var (
	server      *gin.Engine
	ps          services.PetService
	pc          controllers.PetController
	ctx         context.Context
	petc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)
func hello(c *gin.Context) {                             // gin.Context parameter.
	c.IndentedJSON(http.StatusOK,gin.H{"message": "success"})
	}
	
func main() {
    ctx = context.TODO()
	utils.InitializeLogger()
	Logger := utils.Logger

	config.LoadAppConfig()

	metricRouter := gin.Default()
	utils.GetMetricsExporter().Use(metricRouter)

	appRouter := gin.Default()
	middleware.LoadCommonMiddleware(appRouter)
	api.LoadRoutes(appRouter)

    dbClient:=driver.InitializeDbConnection()
	petc = dbClient.Database("petdb").Collection("pets")
	ps = services.NewPetService(petc, ctx)
	pc = controllers.New(ps)
	go func() {
		Logger.Info("App metrics exported on port", zap.String("Port", config.AppConfig.Server.MetricsPort))
		if bootErr := http.ListenAndServe("0.0.0.0:"+config.AppConfig.Server.MetricsPort, metricRouter); bootErr != nil {
			log.Printf("failed to metrics server %v", bootErr)
		}
	}()
	appRouter.GET("/",hello)
	appRouter.POST("/pet/", pc.CreatePet)
	appRouter.GET("/pet/:id", pc.GetPet)
	appRouter.GET("/pet/getall", pc.GetAll)
	appRouter.PUT("/pet", pc.UpdatePet)
	appRouter.DELETE("/pet/:id", pc.DeletePet)
	appRouter.Run(":8000")	

}
