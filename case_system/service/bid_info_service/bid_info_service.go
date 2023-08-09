package bid_info_service

import (
	"case_system/dao/bid_info_dao"

	"github.com/pkg/errors"
)

func GetUserBidInfo(userId string) ([]bid_info_dao.BidInfo, error) {
	bidInfos, err := bid_info_dao.FindByBidderId(userId)
	if err != nil {
		return nil, errors.Wrap(err, "err occur when get BidInfo")
	}

	return bidInfos, nil
}
