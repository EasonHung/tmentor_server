package res

import "mentor/classroom/dao/classroom_info_dao"

type GetClassSettingRes struct {
	ClassSettingList []GetClassSettingResItem `json:"classSettingList"`
}

type GetClassSettingResItem struct {
	SettingName string `json:"settingName"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	ClassTime   int    `json:"classTime"`
	ClassPoints int    `json:"classPoints"`
}

func (this *GetClassSettingRes) ConvertFromClassSettingDao(daoList []classroom_info_dao.ClassSetting) {
	classSettingListVo := make([]GetClassSettingResItem, 0)

	for _, classSettingDao := range daoList {
		classSettingVo := GetClassSettingResItem{
			SettingName: classSettingDao.SettingName,
			Title:       classSettingDao.Title,
			Desc:        classSettingDao.Desc,
			ClassTime:   classSettingDao.ClassTime,
			ClassPoints: classSettingDao.ClassPoints,
		}
		classSettingListVo = append(classSettingListVo, classSettingVo)
	}

	this.ClassSettingList = classSettingListVo
}
