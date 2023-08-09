package vo

type RefreshTokenRes struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewRefreshTokenRes(newAccessToken string, newRefreshToken string) *RefreshTokenRes {
	return &RefreshTokenRes{
		AccessToken: newAccessToken,
		RefreshToken: newRefreshToken,
	}
}