package vo

type CreateThirdPartyUserReq struct {
	ThirdPartyId   string `json:"thirdPartyId"`
	ThirdPartyInfo string `json:"thirdPartyInfo"`
}

type CreateUserReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
