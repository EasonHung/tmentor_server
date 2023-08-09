package dto

type UserDTO struct {
	UserId             string
	UserStatus         string
	UserThirdPartyId   string
	UserThirdPartyInfo string
	UserWalletId       string
}

func NewUserDTO(thirdPartyId string, thirdPartyInfo string) UserDTO {
	return UserDTO{
		UserThirdPartyId:   thirdPartyId,
		UserThirdPartyInfo: thirdPartyInfo,
	}
}
