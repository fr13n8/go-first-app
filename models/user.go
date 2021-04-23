package models

import (
	"github.com/gbrlsnchs/jwt/v3"
	"gorm.io/gorm"
)

type User struct {
	Name       string `json:"name"`
	Age        uint   `json:"age"`
	Email      string `json:"email" gorm:"unique"`
	Verified   bool   `json:"verified"`
	Password   string `json:"-"`
	gorm.Model `json:"-"`
}

type SignUpData struct {
	Name            string `json:"name" validate:"required" binding:"required"`
	Email           string `json:"email" validate:"email" binding:"required"`
	Age             uint   `json:"age" binding:"required,numeric"`
	Password        string `json:"password" validate:"min=5,max=32" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" validate:"eqfield=Password,required" binding:"required"`
}

type SignInData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInResponseData struct {
	User  User   `json:"user" binding:"required"`
	Token string `json:"token" binding:"required"`
}

type UpdatedData struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   uint   `json:"age" binding:"required"`
}

type JwtPayload struct {
	jwt.Payload
}
