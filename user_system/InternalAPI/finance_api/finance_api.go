package finance_api

import (
	"bytes"
	"encoding/json"
	"mentor_app/user_system/initialize"
	"net/http"

	"github.com/pkg/errors"
)

type InitWalletRes struct {
	WalletId string `json:"walletId"`
}

func InitWallet(traceId string, userId string) (error, string) {
	resData := map[string]string{}
	values := map[string]string{}
	values["userId"] = userId
	json_data, err := json.Marshal(values)
	if err != nil {
		err = errors.Wrap(err, "error call init wallet api")
		return err, ""
	}

	url := initialize.GLOBAL_CONFIG.Servers.FinanceSystem + "/wallet/new"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		err = errors.Wrap(err, "error call init wallet api")
		return err, ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("request-id", traceId)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = errors.Wrap(err, "error call init wallet api")
		return err, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.Wrap(err, "error call init wallet api")
		return err, ""
	}

	err = json.NewDecoder(resp.Body).Decode(&resData)
	if err != nil {
		err = errors.Wrap(err, "error call init wallet api")
		return err, ""
	}

	return nil, resData["walletId"]
}
