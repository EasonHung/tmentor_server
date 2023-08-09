package ownership_dao

type OwnerShip struct {
	Id                string   `bson:"_id,omitempty"`
	UserId            string   `bson:"userId"`
	ClassroomId       string   `bson:"classroomId"`
	TeacherClassrooms []string `bson:"teacherClassrooms"`
}
