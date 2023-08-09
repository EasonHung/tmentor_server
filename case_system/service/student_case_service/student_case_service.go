package student_case_service

import (
	"case_system/dao/bid_info_dao"
	"case_system/dao/db_connection"
	"case_system/dao/student_case_dao"
	"case_system/dto"
	"case_system/internalAPI/user_system_api"
	"context"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

const PAGE_PER_SIZE int64 = 30

func AddStudentCase(avatar string, userId string, nickname string, postTime time.Time, title string, content string, pictureUrl []string) error {
	studentCaseId := bson.NewObjectId().Hex()
	bidInfoIds := make([]string, 0)
	entity := student_case_dao.StudentCase{
		StudentCaseId: studentCaseId,
		BidInfoIds:    bidInfoIds,
		AvatarUrl:     avatar,
		UserId:        userId,
		Nickname:      nickname,
		PostTime:      postTime,
		Title:         title,
		Content:       content,
		PictureUrl:    pictureUrl,
	}

	_, err := student_case_dao.InsertOne(entity)

	return err
}

func GetStudentCaseByPage(page int64) ([]student_case_dao.StudentCase, error) {
	cases, err := student_case_dao.FindWithPagination(page, PAGE_PER_SIZE)
	if err != nil {
		return nil, err
	}

	return cases, nil
}

func GetStudentCaseByCaseId(studentCaseId string) (student_case_dao.StudentCase, error) {
	studentCase, err := student_case_dao.FindByCaseId(studentCaseId)
	if err != nil {
		return studentCase, errors.Wrap(err, "err occur when get one student case from db")
	}

	return studentCase, nil
}

func AddBid(biderId string, price int, studentCaseId string, classTime string) error {
	transaction := func(ctx context.Context) (interface{}, error) {
		bidInfoId := bson.NewObjectId().Hex()
		bidInfo := bid_info_dao.BidInfo{
			BidInfoId:    bidInfoId,
			CaseId:       studentCaseId,
			BidderId:     biderId,
			BidPrice:     price,
			BidClassTime: classTime,
		}
		_, err := bid_info_dao.InsertOneWithTx(ctx, bidInfo)
		if err != nil {
			return nil, err
		}
		err = student_case_dao.PushBidInfoByCaseIdWithTx(ctx, studentCaseId, bidInfoId)
		return nil, err
	}

	_, err := db_connection.MONGO_CLIENT.DoTransaction(context.Background(), transaction)
	return err
}

func GetBidInfos(studentCaseId string) ([]dto.BidInfoDto, error) {
	bidInfoIds, err := student_case_dao.FindBidInfoIdByCaseId(studentCaseId)
	if err != nil {
		return nil, errors.Wrap(err, "err occur when get bid info ids")
	}

	bidInfoDtos := make([]dto.BidInfoDto, 0)
	for _, bidInfoId := range bidInfoIds {
		bidInfo, err := bid_info_dao.FindByBidInfoId(bidInfoId)
		if err != nil {
			return nil, errors.Wrap(err, "err occur when get bid info id")
		}

		bidInfoDto := dto.BidInfoDto{}
		err, userInfoMap := user_system_api.GetUserInfo(bidInfo.BidderId)
		if err != nil {
			return nil, errors.Wrap(err, "err occur when get userInfo")
		}
		if userInfoMap == nil {
			return nil, errors.Errorf("no userInfo")
		}
		bidInfoDto.UserInfoMapConvertor(userInfoMap)
		bidInfoDto.BidInfoId = bidInfo.BidInfoId
		bidInfoDto.BidderId = bidInfo.BidderId
		bidInfoDto.BidPrice = bidInfo.BidPrice
		bidInfoDto.StudentCaseId = studentCaseId
		bidInfoDto.BidClassTime = bidInfo.BidClassTime

		bidInfoDtos = append(bidInfoDtos, bidInfoDto)
	}

	return bidInfoDtos, nil
}

func GetStudentCasesByUserId(userId string) ([]student_case_dao.StudentCase, error) {
	cases, err := student_case_dao.FindCasesByUserIdDesc(userId)
	if err != nil {
		return nil, err
	}

	return cases, nil
}
