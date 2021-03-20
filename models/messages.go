package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderId   uint   `json:"SenderId"`
	ReceiverId uint   `json:"ReceiverId"`
	Body       string `json:"Body"`
}
