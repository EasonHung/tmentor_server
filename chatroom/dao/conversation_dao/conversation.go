package conversation_dao

type Conversation struct {
	Id             string       `bson:"_id,omitempty"`
	ConversationId string       `bson:"conversationId"`
	Type           int          `bson:"type"` // 0: single, 1: group
	Participants   []string     `bson:"participants"`
	ReadCursor     []ReadCursor `bson:"readCursor"`
}

type ReadCursor struct {
	UserId   string `bson:"userId"`
	DeviceId string `bson:"deviceId"`
	Cursor   string `bson:"cursor"`
}
