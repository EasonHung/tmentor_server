package bid_info_dao

type BidInfo struct {
	BidInfoId    string `bson:"bidInfoId"`
	CaseId       string `bson:"caseId"`
	BidderId     string `bson:"bidderId"`
	BidPrice     int    `bson:"bidPrice"`
	BidClassTime string `bson:"bidClassTime"`
}
