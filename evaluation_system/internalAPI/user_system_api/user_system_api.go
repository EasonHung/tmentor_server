package user_system_api

import (
	"encoding/json"
	"evaluation_system/config"
	"net/http"

	"github.com/pkg/errors"
)

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

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&resData)
		if err != nil {
			err = errors.Wrap(err, "error call get user info")
			return err, resData
		}
		return nil, resData
	}
	if resData.Code != "0000" {
		err = errors.Wrap(errors.Errorf("error get user Info, error message: %s", resData.Message), "error call get user info")
		return err, resData
	}

	return err, resData
}
