package req

type UpdateClassSettingReq struct {
	SettingName string `json:"settingName"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	ClassTime   int    `json:"classTime"`
	ClassPoints int    `json:"classPoints"`
}
