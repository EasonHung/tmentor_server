package dto

type JwtClaim struct {
	UserId     string `json:"userId"`
	ExpireTime string `json:"expireTime"`
}

func (JwtClaim) Valid() error {
	return nil
}
