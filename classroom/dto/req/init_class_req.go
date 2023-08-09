package req

import "encoding/json"

type InitClassReq struct {
	ClassroomId string `json:"classroomId"`
	MentorId    string `json:"mentorId"`
	StudentId   string `json:"studentId"`
	ClassTitle  string `json:"classTitle"`
	ClassDesc   string `json:"classDesc"`
	ClassTime   int    `json:"classTime"`
	Points      int    `json:"points"`
}

func (this *InitClassReq) ParseFromJsonString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), this)
	if err != nil {
		return err
	}

	return nil
}
