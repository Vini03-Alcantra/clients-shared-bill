package service

import (
	"fmt"
	"log"

	"github.com/ClientsSharedBill/src/dto"
	"github.com/ClientsSharedBill/src/models"
	"github.com/ClientsSharedBill/src/repository"
	"github.com/mashingan/smapping"
)

type ClientService interface {
	Update(client dto.ClientUpdateDTO) models.Client
	GetAllClients() []models.Client
	Delete(c models.Client)
	IsAllowedToEdit(clientID string) bool
}

type clientService struct {
	clientRepository repository.ClientRepository
}

func NewClientService(clientRepo repository.ClientRepository) ClientService {
	return &clientService{
		clientRepository: clientRepo,
	}
}

func (service *clientService) Update(client dto.ClientUpdateDTO) models.Client {
	clientToUpdate := models.Client{}
	err := smapping.FillStruct(&clientToUpdate, smapping.MapFields(&service))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updateClient := service.clientRepository.UpdateClient(clientToUpdate)
	return updateClient
}

func (service *clientService) GetAllClients() []models.Client {
	listClients := service.clientRepository.GetAllClients()
	return listClients
}

func (service *clientService) Delete(c models.Client) {
	service.clientRepository.DeleteClients(c)
}

func (service *clientService) IsAllowedToEdit(clientID string) bool {
	c := service.clientRepository.FindByID(clientID)
	id := fmt.Sprintf("%v", c.ID)
	return clientID == id
}
