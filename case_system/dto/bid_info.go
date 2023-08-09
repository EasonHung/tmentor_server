package dto

type BidInfoDto struct {
	BidInfoId       string
	BidderId        string
	StudentCaseId   string
	BidderAvatorUrl string
	BidderNickname  string
	BidderJob       string
	BidderEducation string
	BidderGender    string
	BidClassTime    string
	BidPrice        int
}

func (this *BidInfoDto) UserInfoMapConvertor(userInfoMap map[string]interface{}) {
	this.BidderNickname = userInfoMap["Nickname"].(string)
	this.BidderAvatorUrl = userInfoMap["AvatorUrl"].(string)
	this.BidderJob = userInfoMap["Profession"].(string)
	this.BidderEducation = userInfoMap["Education"].(string)
	this.BidderGender = userInfoMap["Gender"].(string)
}
