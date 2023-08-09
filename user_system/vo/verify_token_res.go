package vo

import (
	"mentor_app/user_system/db/dao/user_info_dao"
)

type VerifyTokenRes struct {
	UserId       string `json:"userId"`
	UserStatus   string `json:"userStatus"`
	LoginInfoId  int    `json:"loginInfoId"`
	ThirdPartyId string `json:"thirdPartyId"`
}

func (this *VerifyTokenRes) UserDaoConvertor(userInfo user_info_dao.UserInfo) {
	this.UserId = userInfo.UserId
	this.UserStatus = userInfo.UserStatus
}
