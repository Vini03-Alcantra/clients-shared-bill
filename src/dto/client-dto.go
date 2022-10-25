package dto

type ClientUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Cpf      string `json:"cpf" form:"cpf" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required"`
}
