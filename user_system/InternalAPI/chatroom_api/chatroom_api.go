package chatroom_api

import (
	"bytes"
	"encoding/json"
	"mentor_app/user_system/initialize"
	"net/http"

	"github.com/pkg/errors"
)

func InitChatroom(traceId string, userId string) error {
	values := map[string]string{"userId": userId}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	url := initialize.GLOBAL_CONFIG.Servers.Chatroom + "/info/user/init"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		err = errors.Wrap(err, "error call init chatroom api")
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("request-id", traceId)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = errors.Wrap(err, "error call init chatroom api")
		return err
	}

	if resp.StatusCode != 200 {
		return errors.Errorf("error call init chatroom api")
	}

	defer resp.Body.Close()

	return nil
}
