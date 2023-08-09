package user_service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mentor_app/user_system/InternalAPI/chatroom_api"
	"mentor_app/user_system/InternalAPI/classroom_api"
	"mentor_app/user_system/InternalAPI/evaluation_api"
	"mentor_app/user_system/dto"
	Entity "mentor_app/user_system/entity"
	"mentor_app/user_system/errcode"
	"mentor_app/user_system/initialize"
	"mentor_app/user_system/mentor_redis"
	"mentor_app/user_system/middleware/log"
	"mentor_app/user_system/service/service_constant"
	DateUtils "mentor_app/user_system/utils"
	"mentor_app/user_system/utils/cloud_storage_utils"
	"mentor_app/user_system/utils/slice_utils"
	"mentor_app/user_system/vo"
	"mime/multipart"
	"net/http"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"mentor_app/user_system/db/connection"
	"mentor_app/user_system/db/dao/profession_catelog_dao"
	"mentor_app/user_system/db/dao/user_dao"
	"mentor_app/user_system/db/dao/user_info_dao"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

var jwtSecret = []byte(initialize.GLOBAL_CONFIG.Cert.JwtSecret)
var passwordEncryptKey = initialize.GLOBAL_CONFIG.Cert.UserAuthSecret

func CreateUserWithThirdParty(userDTO dto.UserDTO, traceId string) (dto.UserDTO, error) {
	const INITIAL_STATUS string = user_dao.USER_INIT

	transaction := func(tx context.Context) (interface{}, error) {
		newUserId := xid.New().String()

		newThirdPartyInfo := user_info_dao.ThirdParty{
			ThirdPartyId:   userDTO.UserThirdPartyId,
			ThirdPartyInfo: userDTO.UserThirdPartyInfo,
		}

		newUserInfo := user_info_dao.UserInfo{
			UserId:     newUserId,
			UserStatus: INITIAL_STATUS,
			ThirdParty: newThirdPartyInfo,
		}
		user_info_dao.CreateOneWithTx(tx, newUserInfo)

		err := chatroom_api.InitChatroom(traceId, newUserId)
		if err != nil {
			err = errors.Wrap(err, "error occurs when call init chatroom api")
			return nil, err
		}
		err = classroom_api.InitClassroom(traceId, newUserId)
		if err != nil {
			err = errors.Wrap(err, "error occurs when call init classroom api")
			return nil, err
		}
		err = evaluation_api.InitEvaluation(traceId, newUserId)
		if err != nil {
			err = errors.Wrap(err, "error occurs when call init evaluation api")
			return nil, err
		}

		userDTO.UserId = newUserId
		userDTO.UserStatus = INITIAL_STATUS
		return userDTO, nil
	}

	newUserDTO, err := connection.MONGO_CLIENT.DoTransaction(context.TODO(), transaction)
	if err != nil {
		return userDTO, err
	}

	return newUserDTO.(dto.UserDTO), nil
}

func UserLoginWithThirdParty(thirdPartyId string, thirdPartyAccessToken string) (string, string, string, string, error) {
	err := verifyAccessToken(thirdPartyAccessToken)
	if err != nil {
		return "", "", "", "", err
	}

	userEntity, err := user_info_dao.FindByThirdPartyId(thirdPartyId)
	if err != nil {
		log.Logger.Error(err)
		return "", "", "", "", err
	}

	accessToken := createAccessJwt(userEntity.UserId)
	refreshToken := createRefreshJwt(userEntity.UserId)

	return accessToken, refreshToken, userEntity.UserId, userEntity.UserStatus, nil
}

func RefreshToken(RefreshToken string) (string, string, error) {
	claim, err := getClaimFromJwt(RefreshToken)
	if err != nil {
		return "", "", errors.Wrap(err, "token decode error")
	}

	err = checkUserBlackList(claim.UserId)
	if err != nil {
		return "", "", err
	}

	err = checkExpireTime(claim.ExpireTime)
	if err != nil {
		return "", "", err
	}

	newAccessToken := createAccessJwt(claim.UserId)
	newRefreshToken := createRefreshJwt(claim.UserId)
	return newAccessToken, newRefreshToken, nil
}

func GetUserInfo(userId string) (user_info_dao.UserInfo, error) {
	result, err := user_info_dao.FindOneByUserId(userId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func UpdateUserPicture(userId string, pictureFile *multipart.FileHeader) error {
	pictureUrl, err := cloud_storage_utils.UploadAndGetPublicUrl(initialize.GLOBAL_CONFIG.Gcp.BucketName, "user_system/picture/", pictureFile)
	if err != nil {
		return err
	}

	err = user_info_dao.UpdatePictureByUserId(userId, pictureUrl)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserAvator(userId string, avatorFile *multipart.FileHeader) error {
	avatorUrl, err := cloud_storage_utils.UploadAndGetPublicUrl(initialize.GLOBAL_CONFIG.Gcp.BucketName, "user_system/avator/", avatorFile)
	if err != nil {
		return err
	}

	err = user_info_dao.UpdateAvatorUrlByUserId(userId, avatorUrl)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserInfo(userId string, userInfoEntity user_info_dao.UserInfo, userStatus string) error {
	// err, categories := getCategories(userInfoEntity.Profession)
	// if err != nil {
	// 	errors.Wrap(err, "error get catelogs")
	// 	return err
	// }

	// userInfoEntity.ProfessionCategories = categories
	userInfoEntity.UserStatus = userStatus
	err := user_info_dao.UpdateByUserId(userId, userInfoEntity)
	if err != nil {
		return errors.Wrap(err, "error update user info")
	}

	return nil
}

func GetCardsByPage(traceId string, page int64, gender []string, fields []string) ([]vo.GetCardsRes, error) {
	cards, err := user_info_dao.FindByConditionsWithPagination(page, fields, gender)
	if err != nil {
		err = errors.Wrap(err, "error get cards")
		return nil, err
	}

	cards = shuffleCards(cards)

	resVo := convertUserInfoToCardsVo(cards)

	err = fillInStudentCount(traceId, resVo)
	if err != nil {
		return nil, err
	}

	return resVo, nil
}

func GetUser(userId string) (user_info_dao.UserInfo, error) {
	user, err := user_info_dao.FindOneByUserId(userId)
	if err != nil {
		err = errors.Wrap(err, "error get user")
		return user_info_dao.UserInfo{}, err
	}

	return user, nil
}

func GetFcmTokenAndNickname(userId string) (dto.PushNotificationInfoDto, error) {
	pushNotificationInfoDto, err := user_info_dao.FindFcmAndNicknameByUserId(userId)
	if err != nil {
		return dto.PushNotificationInfoDto{}, errors.Wrap(err, "error get user fcm")
	}
	return pushNotificationInfoDto, nil
}

func InsertFcmToken(userId string, fcmToken string) error {
	err := user_info_dao.PushFcmByUserId(userId, fcmToken)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFcmToken(userId string, fcmToken string) error {
	err := user_info_dao.PullFcmByUserId(userId, fcmToken)
	if err != nil {
		return errors.Wrap(err, "error delete fcm token")
	}

	return nil
}

func UpdateFcmToken(userId string, originFcmToken string, newFcmToken string) error {
	err := user_info_dao.PullFcmByUserId(userId, originFcmToken)
	if err != nil {
		return err
	}
	err = user_info_dao.PushFcmByUserId(userId, newFcmToken)
	if err != nil {
		return err
	}
	return nil
}

func createAccessJwt(userId string) string {
	expireTime := DateUtils.ShiftMinutes(service_constant.ACCESS_TOKEN_EXPIRE_TIME).UTC()
	claim := Entity.JwtClaim{
		UserId:     userId,
		ExpireTime: expireTime.Format("2006/1/2 15:04:05"),
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		panic(errcode.TokenCreateError)
	}

	return token
}

func createRefreshJwt(userId string) string {
	expireTime := DateUtils.ShiftDays(service_constant.REFRESH_TOKEN_EXPIRE_TIME).UTC()
	claim := Entity.JwtClaim{
		UserId:     userId,
		ExpireTime: expireTime.Format("2006/1/2 15:04:05"),
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		panic(errcode.TokenCreateError)
	}

	return token
}

func checkExpireTime(tokenTime string) error {
	format := "2006/1/2 15:04:05"
	exprieTime, _ := time.Parse(format, tokenTime)

	if exprieTime.Before(time.Now().UTC()) {
		return errors.Wrap(errors.New("token expired"), "token expired")
	}

	return nil
}

func getClaimFromJwt(token string) (*Entity.JwtClaim, error) {
	var claim *Entity.JwtClaim
	tokenClaims, err := jwt.ParseWithClaims(token, &Entity.JwtClaim{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		err = errors.Wrap(err, "error get token claim")
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Entity.JwtClaim); ok {
		claim = claims
	}

	return claim, nil
}

func verifyAccessToken(thirdPartyAccessToken string) error {
	res, _ := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=" + thirdPartyAccessToken)
	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	response := dto.GoogleTokenVerifyRes{}
	err = json.Unmarshal(sitemap, &response)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if strings.Split(response.Aud, "-")[0] != initialize.GLOBAL_CONFIG.Cert.GoogleAuthClientID &&
		strings.Split(response.Aud, "-")[0] != initialize.GLOBAL_CONFIG.Cert.WebGoogleAuthClientID {
		return errcode.NewError("0001", "google access token error")
	}

	return nil
}

// if don't get catelogs, than insert empty one
func getCategories(professionNames []string) (error, []string) {
	allCatelogs := make([]string, 0)
	for _, professionName := range professionNames {
		catelogs, err := profession_catelog_dao.FindByName(professionName)
		if err != nil {
			errors.Wrap(err, "error occurs when get catelogs")
			return err, nil
		}

		if len(catelogs) == 0 {
			catelog := profession_catelog_dao.ProfessionCategory{
				Name:     professionName,
				Category: "",
			}
			_, err := profession_catelog_dao.CreateOne(catelog)
			if err != nil {
				errors.Wrap(err, "error occurs when insert catelogs")
				return err, nil
			}
		}

		for _, catelog := range catelogs {
			allCatelogs = append(allCatelogs, catelog.Category)
		}
	}

	return nil, allCatelogs
}

func shuffleCards(cards []user_info_dao.UserInfo) []user_info_dao.UserInfo {
	return convInterfaceToCards(slice_utils.RandShuffle(cards))
}

func convInterfaceToCards(interfaceSlice []interface{}) []user_info_dao.UserInfo {
	res := make([]user_info_dao.UserInfo, len(interfaceSlice))

	for index, value := range interfaceSlice {
		res[index] = value.(user_info_dao.UserInfo)
	}
	return res
}

func checkAccountPasswordFormat(account string, password string) error {
	_, err := mail.ParseAddress(account)
	if err != nil {
		return errors.Errorf("帳號格式不正確")
	}

	match, _ := regexp.MatchString("^[A-Za-z0-9]*$", password)
	if match != true {
		return errors.Errorf("密碼格式不正確")
	}
	return nil
}

func checkUserBlackList(userId string) error {
	exist, _ := mentor_redis.Client.SIsMember(service_constant.USER_BLACKLIST, userId).Result()
	if exist {
		return errors.Wrap(errors.New("user is in the black list"), "userToken is in the black list")
	}
	return nil
}

func convertUserInfoToCardsVo(userInfoDaoList []user_info_dao.UserInfo) []vo.GetCardsRes {
	res := make([]vo.GetCardsRes, 0)

	for _, dao := range userInfoDaoList {
		jobExperienceList := make([]vo.JobExperienceVo, 0)
		educationList := make([]vo.EducationVo, 0)

		for _, jobExperienceDao := range dao.JobExperiences {
			jobExperienceVo := vo.JobExperienceVo{
				CompanyName: jobExperienceDao.CompanyName,
				JobName:     jobExperienceDao.JobName,
				StartTime:   jobExperienceDao.StartTime,
				EndTime:     jobExperienceDao.EndTime,
			}
			jobExperienceList = append(jobExperienceList, jobExperienceVo)
		}

		for _, educationDao := range dao.Education {
			educationVo := vo.EducationVo{
				SchoolName: educationDao.SchoolName,
				Subject:    educationDao.Subject,
				StartTime:  educationDao.StartTime,
				EndTime:    educationDao.EndTime,
			}
			educationList = append(educationList, educationVo)
		}

		cardRes := vo.GetCardsRes{
			AvatorUrl:      dao.AvatorUrl,
			UserId:         dao.UserId,
			Nickname:       dao.Nickname,
			AboutMe:        dao.AboutMe,
			Education:      educationList,
			Gender:         dao.Gender,
			Profession:     dao.Profession,
			Fields:         dao.Fields,
			JobExperiences: jobExperienceList,
			PictureUrl:     dao.PictureUrl,
			MentorSkill:    dao.MentorSkill,
			CreateTime:     dao.CreateTime,
			ModifyTime:     dao.ModifyTime,
		}
		res = append(res, cardRes)
	}

	return res
}

func fillInStudentCount(traceId string, cards []vo.GetCardsRes) error {
	for index, card := range cards {
		err, res := classroom_api.GetStudentCount(traceId, card.UserId)
		if err != nil {
			return errors.Wrap(err, "error get card student count")
		}

		cards[index].StudentCount = res.Data.Count
	}

	return nil
}
