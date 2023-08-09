package vo

type AddStudentCaseReq struct {
	AvatarUrl  string `json:"avatarUrl"`
	UserId     string `json:"userId"`
	Nickname   string `json:"nickname"`
	PostTime   string `json:"postTime"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	PictureUrl string `json:"pictureUrl"`
}
