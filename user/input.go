package user

import (
	"BWA/rpcp"
)

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

func (i *RegisterUserInput) FromProto(in *rpcp.RegisterUserInput) *RegisterUserInput {
	i.Name = in.Name
	i.Occupation = in.Occupation
	i.Email = in.Occupation
	i.Password = in.Password
	return i
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
