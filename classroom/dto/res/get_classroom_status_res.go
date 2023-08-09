package res

import classroom_dto "mentor/classroom/domain/classroom/dto"

type GetClassroomStatusRes struct {
	Status string `json:"status"`
	ClassSettingInfo classroom_dto.ClassSetting `json:"classSettingInfo"`
}

func NewGetClassroomStatusRes(status string, classSetting classroom_dto.ClassSetting) GetClassroomStatusRes {
	return GetClassroomStatusRes{
		Status: status,
		ClassSettingInfo: classSetting,
	}
}
