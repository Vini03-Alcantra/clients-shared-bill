package dto

type ClientUpdateDTO struct {
	ID       string `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	CPF      string `json:"cpf" form:"cpf" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required"`
}
