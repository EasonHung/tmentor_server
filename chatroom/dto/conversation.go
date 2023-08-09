package dto

type Conversation struct {
	UserId       string   `json:"userId"`
	Type         int      `json:"type"` // 0: single, 1: group
	Participants []string `json:"participants"`
}
