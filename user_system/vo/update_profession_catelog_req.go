package vo

type UpdateProfssionCatelogReq struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	OldCategory string `json:"oldCategory"`
}
