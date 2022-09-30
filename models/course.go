package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID        string     `gorm:"primary_key" json:"id"`
	Name      string     `db:"name" json:"name"`
	CPF       string     `db:"area" json:"area"`
	Email     string     `json:"teacher"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

func (x *Client) FillDefaults() {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&Client{})
	seed(db)
}

func seed(db *gorm.DB) {
	clients := []Client{
		{ID: "1w2e3r", Name: "João Dória", CPF: "11122233345", Email: "joaodoria@email.com"},
	}
	for _, c := range clients {
		db.Create(&c)
	}
}
