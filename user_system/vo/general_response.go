package vo

import "mentor_app/user_system/errcode"

type GeneralResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *GeneralResponse {
	return &GeneralResponse{
		Code:    "0000",
		Message: "OK",
		Data:    data,
	}
}

func NewErrorResponse(err *errcode.Error) *GeneralResponse {
	return &GeneralResponse{
		Code:    err.Code,
		Message: err.Msg,
	}
}
