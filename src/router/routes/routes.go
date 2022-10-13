package routes

import (
	"net/http"

	client "github.com/ClientsSharedBill/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/clients", client.GetCloentsAll)
	router.GET("/api/v1/clients/:id", client.GetClient)
	router.POST("/api/v1/clients", client.PostClient)
	router.DELETE("/api/v1/clients/:id", client.DeleteClient)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": "Hello world",
		})
	})
	return router
}
