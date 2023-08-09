package ownership_dao

type Ownership struct {
	Id               string         `bson:"_id,omitempty"`
	UserId           string         `bson:"userId"` // 0: single, 1: group
	ConversationList []Conversation `bson:"conversationList"`
}
