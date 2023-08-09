package vo

type NewPostReq struct {
	FromUserId  string `json:"fromUserId"`
	ToUserId    string `json:"toUserId"`
	Score       int    `json:"score"`
	Description string `json:"description"`
}
