package vo

type PayClassBillReq struct {
	ClassId       string `json:"classId"`
	CostSPoint    int    `json:"costSPoint"`
	TeacherUserId string `json:"teacherUserId"`
	StudentUserId string `json:"studentUserId"`
}
