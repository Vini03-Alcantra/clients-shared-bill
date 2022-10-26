package models

import "github.com/google/uuid"

type Client struct {
	ID        string `json:"id" gorm:"primary_key" `
	Name      string `gorm:"type:varchar(255)" json:"name" `
	CPF       string `gorm:"uniqueIndex;type:varchar(255)" json:"cpf"`
	Email     string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string `gorm:"->;<-;not null" json:"-"`
	Token     string `gorm:"-" json:"token,omitempty"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func FillDefaults(client Client) Client {
	if client.ID == "" {
		client.ID = uuid.New().String()
	}

	return client
}
