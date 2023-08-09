package ownership

type Ownership struct {
	Id            string          `bson:"_id,omitempty"`
	UserId        string          `bson:"userId"`
	ClassroomList []ClassroomInfo `bson:"classroomList"`
	StudentList   []string        `bson:"studentList"`
}

type ClassroomInfo struct {
	MentorId    string `bson:"mentorId"`
	ClassroomId string `bson:"classroomId"`
}

func NewOwnership(userId string) Ownership {
	newClassroomList := make([]ClassroomInfo, 0)
	newStudentList := make([]string, 0)

	return Ownership{
		UserId:        userId,
		ClassroomList: newClassroomList,
		StudentList:   newStudentList,
	}
}
