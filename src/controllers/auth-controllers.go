package controllers

import (
	"net/http"

	"github.com/ClientsSharedBill/src/dto"
	helpers "github.com/ClientsSharedBill/src/helper"
	"github.com/ClientsSharedBill/src/models"
	"github.com/ClientsSharedBill/src/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errorDTO := ctx.ShouldBind(&loginDTO)
	if errorDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errorDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredentials(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(models.Client); ok {
		generatedToken := c.jwtService.GenerateToken(v.ID)
		v.Token = generatedToken
		response := helpers.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helpers.BuildErrorResponse("Please check again your credential", "Invalid Credential", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO

	errorDTO := ctx.ShouldBind(&registerDTO)
	if errorDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errorDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdClient := c.authService.CreateClient(registerDTO)
		token := c.jwtService.GenerateToken(createdClient.ID)
		createdClient.Token = token
		response := helpers.BuildResponse(true, "OK!", createdClient)
		ctx.JSON(http.StatusCreated, response)
	}
}
