package dto

import "time"

type Message struct {
	MessageId       string    `json:"messageId"`
	Cmd             int       `json:"cmd"`
	SenderId        string    `json:"senderId"`
	SenderAvatarUrl string    `json:"senderAvatarUrl"`
	RecieverId      string    `json:"recieverId"`
	ConversationId  string    `json:"conversationId"`
	Time            time.Time `json:"time"`
	Message         string    `json:"message"`
	Picture         string    `json:"picture"` // 預覽圖片
	Url             string    `json:"url"`
}
