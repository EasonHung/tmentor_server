package chatroom_api

import (
	"bytes"
	"encoding/json"
	"mentor_app/finance_system/initialize"
	"net/http"
)

func InitChatroom(userId string) error {
	values := map[string]string{"userId": userId}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	url := initialize.GLOBAL_CONFIG.Servers.Chatroom + "/info/user/init"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
