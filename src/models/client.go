package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ID        string         `gorm:"primary_key:auto_increment" json:"id"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	CPF       string         `gorm:"uniqueIndex;type:varchar(255)" json:"cpf"`
	Email     string         `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string         `gorm:"->;<-;not null" json:"password" `
	Token     string         `gorm:"->;<-;not null" json:"-"`
	CreatedAt int64          `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (x *Client) FillDefaults() {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}
}
