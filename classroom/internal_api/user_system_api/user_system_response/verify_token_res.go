package user_system_response

type VerifyTokenRes struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Data    VerifyTokenResData `json:"data"`
}

type VerifyTokenResData struct {
	UserId       string `json:"userId"`
	UserStatus   string `json:"userStatus"`
	LoginInfoId  int    `json:"loginInfoId"`
	ThirdPartyId string `json:"thirdPartyId"`
	WalletId     string `json:"walletId"`
}