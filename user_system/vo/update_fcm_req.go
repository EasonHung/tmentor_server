package vo

type UpdateFcmtokenReq struct {
	UserId         string `json:"userId"`
	OriginFcmToken string `json:"originFcmToken"`
	NewFcmToken    string `json:"newFcmToken"`
}
