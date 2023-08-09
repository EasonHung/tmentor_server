package finance_system_request

type PayClassBillReq struct {
	ClassId       string `json:"classId"`
	CostSPoint    int    `json:"costSPoint"`
	TeacherUserId string `json:"teacherUserId"`
	StudentUserId string `json:"studentUserId"`
}

func NewPayClassBillReq(classId string, costSPoint int, mentorId string, studentId string) PayClassBillReq {
	return PayClassBillReq{
		ClassId: classId,
		CostSPoint: costSPoint,
		TeacherUserId: mentorId,
		StudentUserId: studentId,
	}
}