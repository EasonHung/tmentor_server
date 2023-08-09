package ownership_dao

type Conversation struct {
	ConversationId string   `bson:"conversationId"`
	Type           int      `bson:"type"` // 0: single, 1: group
	Participants   []string `bson:"participants"`
}
