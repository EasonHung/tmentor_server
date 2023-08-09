package vo

type DeleteFcmtokenReq struct {
	UserId   string `json:"userId"`
	FcmToken string `json:"fcmToken"`
}
