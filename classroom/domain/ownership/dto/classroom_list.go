package ownership_dto

type ClassroomList struct {
	ClassroomList []ClassroomInfo `bson:"classroomList"`
}

type ClassroomInfo struct {
	MentorId    string `bson:"mentorId"`
	ClassroomId string `bson:"classroomId"`
}
