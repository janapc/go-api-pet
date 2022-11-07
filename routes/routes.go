package routes

import (
	"go-api-pet/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/api/v1/pets/health", controllers.HealthCheck)
	r.GET("/api/v1/pets", controllers.ListAll)
	r.GET("/api/v1/pets/:id", controllers.FindOneById)
	r.POST("/api/v1/pets", controllers.Add)
	r.DELETE("/api/v1/pets/:id", controllers.Remove)
	r.PATCH("/api/v1/pets/:id", controllers.Update)
	r.GET("/api/v1/pets/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000")
}
