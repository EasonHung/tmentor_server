package res

import (
	"encoding/json"
	"mentor/classroom/dao/class_info_dao"
)

type ClassInfo struct {
	ClassId     string `json:"classId"`
	ClassroomId string `json:"classroomId"`
	Status      string `json:"status"`
	Points      int    `json:"points"`
	ClassTime   int    `json:"classTime"`
	RemainTime  int    `json:"remainTime"`
}

func (this *ClassInfo) ClassInfoDaoConvertor(dao class_info_dao.ClassInfo) {
	this.ClassId = dao.Id
	this.ClassroomId = dao.ClassroomId
	this.Status = dao.Status
	this.Points = dao.Points
	this.ClassTime = dao.ClassTime
	this.RemainTime = dao.ClassTime
}

func (this *ClassInfo) ParseFromJsonString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), this)
	if err != nil {
		return err
	}

	return nil
}
