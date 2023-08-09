package req

type Enroll struct {
	ClassroomToken string `json:"classroomToken"` // 0: single, 1: group
	UserId         string `json:"userId"`
}
