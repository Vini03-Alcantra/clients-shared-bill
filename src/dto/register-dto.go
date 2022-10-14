package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Cpf      string `json:"cpf" form:"cpf" binding:"required"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
