package vo

type LoginRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserId       string `json:"userId"`
	UserStatus   string `json:"userStatus"`
}

func NewLoginRes(accessToken string, refreshToken string, userId string, userStatus string) *LoginRes {
	return &LoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       userId,
		UserStatus:   userStatus,
	}
}
