package res

type InitClassRes struct {
	ClassId string `json:"classId"`
}

func NewInitClassRes(classId string) InitClassRes {
	return InitClassRes{
		ClassId: classId,
	}
}
