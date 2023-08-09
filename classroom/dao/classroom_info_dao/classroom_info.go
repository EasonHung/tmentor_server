package classroom_info_dao

type ClassroomInfo struct {
	Id                string         `bson:"_id,omitempty"`
	ClassroomId       string         `bson:"classroomId"`
	OwnerUserId       string         `bson:"ownerUserId"`
	Students          []string       `bson:"students"`
	Title             string         `bson:"title"`
	TeacherClassrooms []string       `bson:"teacherClassrooms"`
	ClassSettingList  []ClassSetting `bson:"classSettingList"`
}

type ClassSetting struct {
	SettingName string `bson:"settingName"`
	Title       string `bson:"title"`
	Desc        string `bson:"desc"`
	ClassTime   int    `bson:"classTime"`
	ClassPoints int    `bson:"classPoints"`
}
