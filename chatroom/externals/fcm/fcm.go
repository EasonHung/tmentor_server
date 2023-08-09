package fcm

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mentor_app/chatroom/config"
	"net/http"
)

func PostNotificationByToken(token string, title string, body string, conversationId string, senderId string, messageId string, timeStr string, avatarUrl string) (string, error) {
	fcmUrl := "https://fcm.googleapis.com/fcm/send"
	jsonStr := fmt.Sprintf(`{
		"data": {
			"conversationId": "%s",
			"senderId": "%s",
			"nickname": "%s",
			"message": "%s",
			"messageId": "%s",
			"avatarUrl": "%s",
			"time": "%s",
		},
		"to": "%s"
	}`, conversationId, senderId, title, body, messageId, avatarUrl, timeStr, token)

	req, err := http.NewRequest("POST", fcmUrl, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+config.GLOBAL_CONFIG.Fcm.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBody, _ := ioutil.ReadAll(resp.Body)

	return string(resBody), nil
}
