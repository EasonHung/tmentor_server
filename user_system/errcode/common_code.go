package errcode

var (
	TokenDecodeError   = NewError("US-0001", "token decode error")
	UserBlackListError = NewError("US-0002", "user is in the black list")
	TokenCreateError   = NewError("US-0003", "create token error")
	TokenExpiredError  = NewError("US-0004", "token expired")
	ReqBodyError       = NewError("0002", "request body error")
	UnexpectError      = NewError("9999", "unexpect error")
)
