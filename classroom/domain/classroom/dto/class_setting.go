package classroom_dto

import "encoding/json"

type ClassSetting struct {
	SettingName string `json:"settingName" bson:"settingName"`
	Title       string `json:"title" bson:"title"`
	Desc        string `json:"desc" bson:"desc"`
	ClassTime   int    `json:"classTime" bson:"classTime"`
	ClassPoints int    `json:"classPoints" bson:"classPoints"`
}

func (this *ClassSetting) ParseFromJsonString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), this)
	if err != nil {
		return err
	}

	return nil
}
