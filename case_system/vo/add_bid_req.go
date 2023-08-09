package vo

type AddBidReq struct {
	BidPrice      int    `json:"bidPrice"`
	StudentCaseId string `json:"studentCaseId"`
	Classtime     string `json:"classTime"`
}
