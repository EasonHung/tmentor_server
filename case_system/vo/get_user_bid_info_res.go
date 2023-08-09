package vo

import "case_system/dao/bid_info_dao"

type GetUserBidInfoResVo struct {
	Data []GetUserBidInfoResVoItem `json:"data"`
}

type GetUserBidInfoResVoItem struct {
	CaseId       string `json:"caseId"`
	BidClassTime string `json:"bidClassTime"`
	BidInfoId    string `json:"bidInfoId"`
	BidPrice     int    `json:"bidPrice"`
}

func (this *GetUserBidInfoResVo) BidInfoListConvertor(bidInfoList []bid_info_dao.BidInfo) {
	bidInfoListRes := make([]GetUserBidInfoResVoItem, 0)

	for _, bidInfoEntity := range bidInfoList {
		resItem := GetUserBidInfoResVoItem{
			CaseId:       bidInfoEntity.CaseId,
			BidClassTime: bidInfoEntity.BidClassTime,
			BidInfoId:    bidInfoEntity.BidInfoId,
			BidPrice:     bidInfoEntity.BidPrice,
		}
		bidInfoListRes = append(bidInfoListRes, resItem)
	}

	this.Data = bidInfoListRes
}
