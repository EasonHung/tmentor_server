package vo

import "case_system/dto"

type GetBidInfoResVo struct {
	Data []GetBidInfoResVoItem `json:"data"`
}

type GetBidInfoResVoItem struct {
	BidInfoId       string `json:"bidInfoId"`
	BidderId        string `json:"bidderId"`
	StudentCaseId   string `json:"studentCaseId"`
	BidderAvatorUrl string `json:"bidderAvatorUrl"`
	BidderNickname  string `json:"bidderNickname"`
	BidderJob       string `json:"bidderJob"`
	BidderEducation string `json:"bidderEducation"`
	BidderGender    string `json:"bidderGender"`
	BidderClassTime string `json:"bidderClassTime"`
	BidPrice        int    `json:"bidPrice"`
}

func (this *GetBidInfoResVo) DtoConvertor(dtoList []dto.BidInfoDto) {
	voList := make([]GetBidInfoResVoItem, 0)
	for _, dto := range dtoList {
		voItem := GetBidInfoResVoItem{}
		voItem.BidInfoId = dto.BidInfoId
		voItem.BidderId = dto.BidderId
		voItem.StudentCaseId = dto.StudentCaseId
		voItem.BidderAvatorUrl = dto.BidderAvatorUrl
		voItem.BidderNickname = dto.BidderNickname
		voItem.BidderJob = dto.BidderJob
		voItem.BidderEducation = dto.BidderEducation
		voItem.BidderGender = dto.BidderGender
		voItem.BidPrice = dto.BidPrice
		voItem.BidderClassTime = dto.BidClassTime
		voList = append(voList, voItem)
	}
	this.Data = voList
}
