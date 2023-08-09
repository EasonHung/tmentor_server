package vo

import "mentor_app/chatroom/dao/ownership_dao"

type GetConversationListRes struct {
	Data []GetConversationListResItem
}

type GetConversationListResItem struct {
	ConversationId  string     `json:"conversationId"`
	UserInfo        UserInfoVo `json:"userInfo"`
	Type            int        `json:"type"`
	Participants    []string   `json:"participants"`
	UnReadedCount   int        `json:"unReadedCount"`
	LastMessage     string     `json:"lastMessage"`
	LastMessageTime string     `json:"lastMessageTime"`
}

type UserInfoVo struct {
	AvatorUrl string `json:"avatorUrl"`
	UserId    string `json:"userId"`
	Nickname  string `json:"nickname"`
}

func (this *GetConversationListRes) ConversationConvertor(conversationEntityList []ownership_dao.Conversation) {
	conversationList := make([]GetConversationListResItem, 0)

	for _, conversationInfo := range conversationEntityList {
		resItem := GetConversationListResItem{
			ConversationId: conversationInfo.ConversationId,
			Type:           conversationInfo.Type,
			Participants:   conversationInfo.Participants,
		}
		conversationList = append(conversationList, resItem)
	}
	this.Data = conversationList
}
