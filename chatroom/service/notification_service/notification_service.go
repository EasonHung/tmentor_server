package notification_service

import (
	"encoding/json"
	"mentor_app/chatroom/externals/fcm"
	"mentor_app/chatroom/internalAPI/constants/api_response_code"
	"mentor_app/chatroom/internalAPI/user_system_api"
	"mentor_app/chatroom/middleware/log"

	"github.com/pkg/errors"
)

func PushChatNotification(receiverId string, body string, conversationId string, senderId string, messageId string, timeStr string, avatarUrl string) error {
	err, apiRes := user_system_api.GetFcmTokenAndNickname(receiverId)
	if err != nil {
		return err
	}
	if apiRes.Code != api_response_code.SUCESS {
		return errors.New("error get fcm token, message: " + apiRes.Message)
	}
	pushNotificationInfo := apiRes.Data

	// if is loged out then don't send notification
	if len(pushNotificationInfo.FcmToken) == 0 {
		return nil
	}

	for _, fcmToken := range pushNotificationInfo.FcmToken {
		respString, err := fcm.PostNotificationByToken(fcmToken, pushNotificationInfo.Nickname, body, conversationId, senderId, messageId, timeStr, avatarUrl)
		if err != nil {
			return errors.Wrap(err, "error push notification")
		}

		jsonMap := map[string]interface{}{}
		json.Unmarshal([]byte(respString), &jsonMap)
		if jsonMap["success"].(float64) != 1 {
			log.Logger.Warnf("push notification fail, token: %s", fcmToken)
		}
	}

	return nil
}
