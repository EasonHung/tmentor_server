package vo

type UpdateReadCursorReq struct {
	UserId         string `json:"userId"`
	ConversationId string `json:"conversationId"`
	DeviceId       string `json:"deviceId"`
	LastMessageId  string `json:"lastMessageId"`
}
