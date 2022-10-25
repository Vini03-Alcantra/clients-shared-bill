package routes

import (
	"github.com/ClientsSharedBill/src/controllers"
	client "github.com/ClientsSharedBill/src/controllers"
	"github.com/ClientsSharedBill/src/database"
	"github.com/ClientsSharedBill/src/middleware"
	"github.com/ClientsSharedBill/src/repository"
	"github.com/ClientsSharedBill/src/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB                    = database.GetDatabaseConnection()
	clientRepository repository.ClientRepository = repository.NewClientRepository(db)
	jwtService       service.JWTService          = service.NewJWTService()
	authService      service.AuthService         = service.NewAuthService(clientRepository)
	authController   controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := router.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/api/v1/clients", client.GetCloentsAll)
		userRoutes.GET("/api/v1/clients/:id", client.GetClient)
		userRoutes.POST("/api/v1/clients", client.PostClient)
		userRoutes.DELETE("/api/v1/clients/:id", client.DeleteClient)
	}

	return router
}
