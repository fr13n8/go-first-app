package models

import "gorm.io/gorm"

type Message struct {
	SenderId   uint   `json:"SenderId"`
	ReceiverId uint   `json:"ReceiverId"`
	Body       string `json:"Body"`
	gorm.Model
}
