package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ID        string         `gorm:"primary_key" json:"id"`
	Name      string         `db:"name" json:"name"`
	CPF       string         `db:"area" json:"area"`
	Email     string         `json:"teacher"`
	CreatedAt int64          `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (x *Client) FillDefaults() {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}
}

func seed(db *gorm.DB) {
	clients := []Client{
		{ID: "1w2e3r", Name: "João Dória", CPF: "11122233345", Email: "joaodoria@email.com"},
	}
	for _, c := range clients {
		db.Create(&c)
	}
}
