package user_api

import (
	"bytes"
	"encoding/json"
	"io"
	"mentor_app/finance_system/initialize"
	"net/http"

	"github.com/pkg/errors"
)

func VerifyToken(userToken string) (error, map[string]interface{}) {
	resData := map[string]interface{}{}
	values := map[string]string{"userToken": userToken}
	json_data, err := json.Marshal(values)
	if err != nil {
		err = errors.Wrap(err, "error call verify token api")
		return err, nil
	}

	url := initialize.GLOBAL_CONFIG.Servers.UserSystem + "/user/token/verify"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		err = errors.Wrap(err, "error call verify token api")
		return err, nil
	}

	if resp.StatusCode != 200 {
		err = errors.Wrap(err, "error call verify token api")
		return err, nil
	}

	err = json.NewDecoder(resp.Body).Decode(&resData)
	if err != nil {
		err = errors.Wrap(err, "error call verify token api")
		return err, nil
	}
	defer resp.Body.Close()

	return nil, resData
}

func GetWalletId(userId string) (error, string) {
	url := initialize.GLOBAL_CONFIG.Servers.UserSystem + "/user/walletId?"
	resp, err := http.Get(url + "userId=" + userId)
	if err != nil {
		err = errors.Wrap(err, "error call verify token api")
		return err, ""
	}

	if resp.StatusCode != 200 {
		err = errors.Wrap(err, "error call verify token api")
		return err, ""
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return errors.Wrap(err, "err occur when read body bytes"), ""
	}
	res := string(bodyBytes)
	resLen := len(res)
	return nil, res[1 : resLen-1]
}
