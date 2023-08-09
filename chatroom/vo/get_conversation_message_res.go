package vo

import (
	"mentor_app/chatroom/dao/chat_message_dao"
	"mentor_app/chatroom/utils/time_utils"
)

type GetConversationMessageRes struct {
	Data []GetConversationMessageResItem `json:"messages"`
}

type GetConversationMessageResItem struct {
	ConversationId string `json:"conversationId"`
	Cmd            int    `json:"cmd"`
	SenderId       string `json:"senderId"`
	RecieverId     string `json:"recieverId"`
	MessageId      string `json:"messageId"`
	Time           string `json:"time"`
	Message        string `json:"message"`
	Picture        string `json:"picture"`
	Url            string `json:"url"`
}

func (this *GetConversationMessageRes) ChatMessageListConvertor(chatMessageList []*chat_message_dao.ChatMessage) {
	conversationMessagList := make([]GetConversationMessageResItem, 0)

	for _, chatMessage := range chatMessageList {
		time, _ := time_utils.TimeInTaipei(chatMessage.Data.Time)
		resItem := GetConversationMessageResItem{
			ConversationId: chatMessage.ConversationId,
			Cmd:            chatMessage.Data.Cmd,
			SenderId:       chatMessage.SenderId,
			RecieverId:     chatMessage.Data.RecieverId,
			MessageId:      chatMessage.MessageId,
			Time:           time.UTC().Format("2006-01-02T15:04:05Z"),
			Message:        chatMessage.Data.Message,
			Picture:        chatMessage.Data.Picture,
			Url:            chatMessage.Data.Url,
		}
		conversationMessagList = append(conversationMessagList, resItem)
	}
	this.Data = conversationMessagList
}
