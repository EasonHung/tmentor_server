package ws_message

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type ClassSettingInfo struct {
	ClassroomId string `json:"classroomId"`
	SettingName string `json:"settingName"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	ClassTime   int    `json:"classTime"`
	ClassPoints int    `json:"classPoints"`
}

func ParseFromJsonString(jsonString string) (error, ClassSettingInfo) {
	newClassSettingInfo := ClassSettingInfo{}
	err := json.Unmarshal([]byte(jsonString), &newClassSettingInfo)
	if err != nil {
		return errors.WithStack(err), newClassSettingInfo
	}

	return nil, newClassSettingInfo
}

func (this *ClassSettingInfo) ToJsonString() (error, string) {
	classSettingInfoVoByte, err := json.Marshal(this)
	if err != nil {
		return errors.WithStack(err), ""
	}

	return nil, string(classSettingInfoVoByte)
}
