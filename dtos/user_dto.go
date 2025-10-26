package dtos

import "mime/multipart"

// untuk register user baru
type CreateUserInput struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
}

// untuk update profil user
type UpdateUserInput struct {
	Name   string                `form:"name"`
	Email  string                `form:"email"`
	Avatar *multipart.FileHeader `form:"avatar"`
}

// untuk response user ke frontend
type UserResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}
