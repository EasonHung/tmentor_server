package classroom_system_api

import (
	"encoding/json"
	"mentor_app/chatroom/config"
	internal_api "mentor_app/chatroom/internalAPI"
	"mentor_app/chatroom/internalAPI/classroom_system_api/classroom_system_response"
)

var classroomApiHandler internal_api.InternalApiHandler

func init() {
	classroomApiHandler = internal_api.NewApiHandler(config.GLOBAL_CONFIG.ClassroomSystem.BaseUrl)
}

func GetUserClassroomToken(userId string) (error, classroom_system_response.GetClassroomTokenRes) {
	res := classroom_system_response.GetClassroomTokenRes{}

	err, resBody := classroomApiHandler.Get("/internal/info/token?userId=" + userId)
	if err != nil {
		return err, res
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return err, res
	}

	return nil, res
}
