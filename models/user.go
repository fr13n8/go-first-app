package models

import (
	"github.com/gbrlsnchs/jwt/v3"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string    `json:"name"`
	Age              uint      `json:"age"`
	Email            string    `json:"email" gorm:"unique	"`
	Password         string    `json:"password"`
	SendedMessages   []Message `gorm:"foreignKey:SenderId"`
	ReceivedMessages []Message `gorm:"foreignKey:ReceiverId"`
}

type SignUpData struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Age             uint   `json:"age" binding:"required"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type SignInData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdatedData struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   uint   `json:"age" binding:"required"`
}

type JwtPayload struct {
	jwt.Payload
	Email string `json:"email,omitempty"`
	Age   int    `json:"age,omitempty"`
	Name  string `json:"name,omitempty"`
}
