package vo

import "mentor_app/user_system/dto"

type GetFcmTokenAndNicknameRes struct {
	FcmToken []string `json:"fcmToken"`
	Nickname string   `json:"nickname"`
}

func (this *GetFcmTokenAndNicknameRes) DtoConvertor(dto dto.PushNotificationInfoDto) {
	this.FcmToken = dto.FcmToken
	this.Nickname = dto.Nickname
}
