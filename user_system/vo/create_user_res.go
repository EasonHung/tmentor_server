package vo

type CreateUserRes struct {
	Token string `json:"token"`
}

func NewCreateUserRes(token string) *CreateUserRes {
	return &CreateUserRes{
		Token: token,
	}
}
