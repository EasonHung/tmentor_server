package res

import (
	"mentor/classroom/dao/classroom_info_dao"
)

type GetClassroomOwnershipRes struct {
	UserId      string `json:"userId"`
	ClassroomId string `json:"classroomId"`
}

func (this *GetClassroomOwnershipRes) ClassroomInfoConvertor(claassroomInfo classroom_info_dao.ClassroomInfo) {
	this.UserId = claassroomInfo.OwnerUserId
	this.ClassroomId = claassroomInfo.ClassroomId
}
