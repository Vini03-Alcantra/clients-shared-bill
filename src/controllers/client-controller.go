package controllers

import (
	"fmt"
	"net/http"

	"github.com/ClientsSharedBill/src/dto"
	helpers "github.com/ClientsSharedBill/src/helper"
	"github.com/ClientsSharedBill/src/models"
	"github.com/ClientsSharedBill/src/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ClientController interface {
	Update(context *gin.Context)
	GetClientsAll(context *gin.Context)
	DeleteClient(context *gin.Context)
}

type clientController struct {
	clientService service.ClientService
	jwtService    service.JWTService
}

func NewClientController(clientService service.ClientService, jwtService service.JWTService) ClientController {
	return &clientController{
		clientService: clientService,
		jwtService:    jwtService,
	}
}

func (c *clientController) GetClientsAll(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	fmt.Print(token)
	var clients []models.Client = c.clientService.GetAllClients()
	res := helpers.BuildResponse(true, "OK", clients)
	context.JSON(http.StatusOK, res)
}

func (c *clientController) Update(context *gin.Context) {
	var clientUpdateDTO dto.ClientUpdateDTO
	errDTO := context.ShouldBind(&clientUpdateDTO)

	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	clientID := fmt.Sprintf("%v", claims["user_id"])
	if c.clientService.IsAllowedToEdit(clientID) {
		result := c.clientService.Update(clientUpdateDTO)
		res := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, res)
	} else {
		res := helpers.BuildErrorResponse("You dont have any permission", "You aren't the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, res)
	}
}

func (c *clientController) DeleteClient(context *gin.Context) {
	var client models.Client
	id := context.Param("id")
	if id == "" {
		response := helpers.BuildErrorResponse("Failed tou get id", "No param id were found", helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	client.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	clientID := fmt.Sprintf("%v", claims["user_id"])
	if c.clientService.IsAllowedToEdit(clientID) {
		c.clientService.Delete(client)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}
