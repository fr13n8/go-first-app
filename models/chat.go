package models

import "gorm.io/gorm"

type Conversation struct {
	UserID     uint   `json:"userId"`
	ReceiverId uint   `json:"ReceiverId"`
	Type       string `json:"type" gorm:"type:enum('user', 'group');default:'user'"`
	gorm.Model `json:"-"`
}

type NewMessage struct {
	ConversationId uint   `json:"conversationId"`
	Body           string `json:"body"`
}

type Message struct {
	ConversationId uint   `json:"conversationId"`
	Body           string `json:"body"`
	gorm.Model     `json:"-"`
}

type Friend struct {
	OwnerId    uint `json:"ownerId"`
	FriendId   uint `json:"friendId"`
	gorm.Model `json:"-"`
}

type Group struct {
	Name       string `json:"name"`
	gorm.Model `json:"-"`
}

type GroupUser struct {
	UserId     uint `json:"userId"`
	GroupId    uint `json:"groupId"`
	gorm.Model `json:"-"`
}
