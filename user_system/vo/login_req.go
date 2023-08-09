package vo

type LoginWithThirdParyReq struct {
	ThirdPartyId          string `json:"thirdPartyId"`
	ThirdPartyAccessToken string `json:"thirdPartyAccessToken"`
}

type LoginWithAccountReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
