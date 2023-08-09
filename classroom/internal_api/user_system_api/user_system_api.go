package user_system_api

import (
	"encoding/json"
	"mentor/classroom/config"
	"mentor/classroom/internal_api"
	"mentor/classroom/internal_api/user_system_api/user_system_request"
	"mentor/classroom/internal_api/user_system_api/user_system_response"

	"github.com/pkg/errors"
)

var userApiHandler internal_api.InternalApiHandler

func init() {
	userApiHandler = internal_api.NewApiHandler(config.GLOBAL_CONFIG.UserSystem.BaseUrl)
}

func GetUserInfo(userId string) (error, user_system_response.GetUserInfoRes) {
	resData := user_system_response.GetUserInfoRes{}

	err, resBody := userApiHandler.Get("/userInfo/?userId=" + userId)
	if err != nil {
		return err, resData
	}

	err = json.Unmarshal(resBody, &resData)
	if err != nil {
		return errors.WithStack(err), resData
	}

	return nil, resData
}

func VerifyToken(token string) (error, user_system_response.VerifyTokenRes) {
	resData := user_system_response.VerifyTokenRes{}

	reqBody := user_system_request.VerifyTokenReq{
		UserToken: token,
	}
	resBody, err := userApiHandler.Post("/user/token/verify", reqBody)
	if err != nil {
		return err, user_system_response.VerifyTokenRes{}
	}

	err = json.Unmarshal(resBody, &resData)
	if err != nil {
		return errors.WithStack(err), resData
	}

	return nil, resData
}
