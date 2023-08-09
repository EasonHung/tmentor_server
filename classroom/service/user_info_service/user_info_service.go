package user_info_service

import (
	"fmt"
	"mentor/classroom/internal_api/user_system_api"
	"mentor/classroom/internal_api/user_system_api/user_system_error_code"
	"mentor/classroom/internal_api/user_system_api/user_system_response"
	"mentor/classroom/middleware/log"

	"github.com/pkg/errors"
)

func GetUserInfo(userId string) (user_system_response.GetUserInfoResData, error) {
	err, userInfoRes := user_system_api.GetUserInfo(userId)
	if err != nil {
		return user_system_response.GetUserInfoResData{}, err
	}

	if userInfoRes.Code != user_system_error_code.SUCCESS {
		log.Logger.Info(fmt.Sprintf("[user info service] error get user info, userId: %s, status code: %s", userId, userInfoRes.Code))
		return user_system_response.GetUserInfoResData{}, errors.New("error get user info")
	}

	return userInfoRes.Data, nil
}