package routes

import (
	"fmt"
	"net/http"

	"github.com/ShiranaiZo/experiment-golang/app/controllers"
	"github.com/ShiranaiZo/experiment-golang/config"
	responses "github.com/ShiranaiZo/experiment-golang/helpers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, controller controllers.IControllerRegistry) {
	mainEndpoint := fmt.Sprintf("/%s/%s", config.MainConfig.RouteAPIPath, config.MainConfig.RouteAPIVersion)
	group := router.Group(mainEndpoint)

	router.GET("/", func(ctx *gin.Context) {
		responses.HttpResponse(responses.ParamHTTPResponse{
			Code: http.StatusOK,
			Ctx:  ctx,
		})
	})

	group.POST("/users", controller.GetUserController().CreateUser)
	group.GET("/users", controller.GetUserController().GetUsers)
	group.GET("/users/:userId", controller.GetUserController().GetUser)

	// router.GET("/health", healthHandler.HealthCheck)
	// router.GET("/contacts", contactHandler.GetContacts)
	// router.GET("/contacts/:id", contactHandler.GetContact)
	// router.POST("/contacts", contactHandler.CreateContact)
	// router.PUT("/contacts/:id", contactHandler.UpdateContact)
	// router.DELETE("/contacts/:id", contactHandler.DeleteContact)
}
