package repository

import (
	"log"

	"github.com/ClientsSharedBill/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ClientRepository interface {
	InsertClient(client models.Client) models.Client
	UpdateClient(client models.Client) models.Client
	VerifyCredentials(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.Client
}

type clientConnection struct {
	connection *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientConnection{
		connection: db,
	}
}

func (db *clientConnection) InsertClient(client models.Client) models.Client {
	client.Password = hashAndSalt([]byte(client.Password))
	db.connection.Save(&client)
	return client
}

func (db *clientConnection) UpdateClient(client models.Client) models.Client {
	if client.Password != "" {
		client.Password = hashAndSalt([]byte(client.Password))
	} else {
		var tempClient models.Client
		db.connection.Find(&tempClient, client.ID)
		client.Password = tempClient.Password
	}

	db.connection.Save(&client)
	return client
}

func (db *clientConnection) VerifyCredentials(email string, password string) interface{} {
	var client models.Client
	res := db.connection.Where("email = ?", email).Take(&client)
	if res.Error == nil {
		return client
	}

	return nil
}

func (db *clientConnection) FindByEmail(email string) models.Client {
	var client models.Client
	db.connection.Where("email = ?", email).Take(&client)
	return client
}

func (db *clientConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var client models.Client
	return db.connection.Where("email = ?", email).Take(&client)
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}

	return string(hash)
}
