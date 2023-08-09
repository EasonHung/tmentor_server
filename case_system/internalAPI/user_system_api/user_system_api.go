package user_system_api

import (
	"case_system/config"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func GetUserInfo(userId string) (error, map[string]interface{}) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", config.GLOBAL_CONFIG.UserSystem.BaseUrl+"/userInfo/?userId="+userId, nil)
	if err != nil {
		return errors.Wrap(err, "err occur when request get user info"), nil
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "err occur when do get user info"), nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "err occur when read body bytes"), nil
		}
		bodyString := string(bodyBytes)

		err, res := jsonToMap(bodyString)
		return nil, res
	}

	return err, nil
}

func jsonToMap(jsonString string) (error, map[string]interface{}) {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &mapResult)
	if err != nil {
		return errors.Wrap(err, "err occur when unmarshel json"), nil
	}
	return nil, mapResult
}
