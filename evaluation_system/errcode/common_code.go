package errcode

var (
	UserNotFoundError = NewError("ES-0001", "user not found")
	TokenDecodeError  = NewError("US-0001", "token decode error")
	UnexpectError     = NewError("9999", "unexpect error")
)
