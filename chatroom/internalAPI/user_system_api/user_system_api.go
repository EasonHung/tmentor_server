package user_system_api

import (
	"encoding/json"
	"io"
	"mentor_app/chatroom/config"
	"net/http"

	"github.com/pkg/errors"
)

func GetUserAvatorUrl(userId string) (error, string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", config.GLOBAL_CONFIG.UserSystem.BaseUrl+"/userInfo/avatorUrl?userId="+userId, nil)
	if err != nil {
		return errors.Wrap(err, "err occur when request get user info"), ""
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "err occur when do get user info"), ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "err occur when read body bytes"), ""
		}
		res := string(bodyBytes)
		resLen := len(res)
		return nil, res[1 : resLen-1]
	}

	return err, ""
}

func GetUserInfo(userId string) (error, GetUserInfoRes) {
	client := &http.Client{}
	resData := GetUserInfoRes{}

	req, err := http.NewRequest("GET", config.GLOBAL_CONFIG.UserSystem.BaseUrl+"/userInfo/?userId="+userId, nil)
	if err != nil {
		return errors.Wrap(err, "err occur when request get user info"), resData
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "err occur when do get user info"), resData
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, "err occur when read body bytes"), resData
	}

	err = json.NewDecoder(resp.Body).Decode(&resData)
	if err != nil {
		err = errors.Wrap(err, "err occur when do get user info")
		return err, resData
	}

	return err, resData
}

func GetFcmTokenAndNickname(userId string) (error, GetFcmTokenAndNicknameRes) {
	client := &http.Client{}
	resData := GetFcmTokenAndNicknameRes{}

	req, err := http.NewRequest("GET", config.GLOBAL_CONFIG.UserSystem.BaseUrl+"/userInfo/fcmTokenAndNickname?userId="+userId, nil)
	if err != nil {
		return errors.Wrap(err, "err occur when request get user fcm info"), resData
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "err occur when do get user fcm info"), resData
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, "err occur when read body bytes"), resData
	}

	err = json.NewDecoder(resp.Body).Decode(&resData)
	if err != nil {
		err = errors.Wrap(err, "err occur when do get user info")
		return err, resData
	}

	return err, resData
}

func jsonToMap(jsonString string) (error, map[string]interface{}) {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &mapResult)
	if err != nil {
		return errors.Wrap(err, "err occur when unmarshel json"), nil
	}
	return nil, mapResult
}
