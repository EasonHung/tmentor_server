package errcode

import "fmt"

type Error struct {
	Code string
	Msg  string
}

var errorCode = make(map[string]string)

func NewError(code string, msg string) *Error {
	if _, ok := errorCode[code]; ok {
		panic(fmt.Sprintf("錯誤ID->%s已存在，請換一個", code))
	}
	errorCode[code] = msg

	return &Error{
		Code: code,
		Msg:  msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %s, msg: %s", e.Code, e.Msg)
}
