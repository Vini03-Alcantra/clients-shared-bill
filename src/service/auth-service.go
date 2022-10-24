package service

import (
	"log"

	"github.com/ClientsSharedBill/src/dto"
	"github.com/ClientsSharedBill/src/models"
	"github.com/ClientsSharedBill/src/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredentials(email string, password string) interface{}
	CreateClient(client dto.ClientCreateDTO) models.Client
	FindByEmail(email string) models.Client
	IsDuplicateEmail(email string) bool
}

type authService struct {
	clientRepository repository.ClientRepository
}

func NewAuthService(clientRep repository.ClientRepository) AuthService {
	return &authService{
		clientRepository: clientRep,
	}
}

func (service *authService) VerifyCredentials(email string, password string) interface{} {
	res := service.clientRepository.VerifyCredentials(email, password)
	if v, ok := res.(models.Client); ok {
		comparePassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparePassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateClient(client dto.ClientCreateDTO) models.Client {
	clientToCreate := models.Client{}
	err := smapping.FillStruct(&clientToCreate, smapping.MapFields(&client))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.clientRepository.InsertClient(clientToCreate)
	return res
}

func (service *authService) FindByEmail(email string) models.Client {
	return service.clientRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.clientRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
