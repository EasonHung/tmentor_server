package user_info_dao

import (
	"context"
	"mentor_app/user_system/db/connection"
	"mentor_app/user_system/db/dao/user_dao"
	"mentor_app/user_system/dto"
	"time"

	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson"
)

const PAGE_SIZE = 20

var collection *qmgo.Collection

func init() {
	collection = connection.MONGO_CLIENT.Database("user_info_system").Collection("user_info")
}

func CreateOneWithTx(ctx context.Context, obj UserInfo) (*qmgo.InsertOneResult, error) {
	emptyMentorSkillList := make([]string, 0)
	obj.CreateTime = time.Now()
	obj.ModifyTime = time.Now()
	obj.MentorSkill = emptyMentorSkillList
	obj.FcmToken = make([]string, 0)
	result, err := collection.InsertOne(ctx, obj)
	return result, err
}

func FindOneByUserId(userId string) (UserInfo, error) {
	result := UserInfo{}
	err := collection.Find(context.Background(), bson.M{"userId": userId}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func FindByThirdPartyId(thirdPartyId string) (UserInfo, error) {
	result := UserInfo{}
	err := collection.Find(context.Background(), bson.M{"thirdParty.thirdPartyId": thirdPartyId}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func FindByUserIds(userId []string) ([]UserInfo, error) {
	result := []UserInfo{}
	err := collection.Find(context.Background(), bson.M{"userId": bson.M{"$in": userId}}).All(result)
	if err != nil {
		err = errors.Wrap(err, "error find userInfos")
		return result, err
	}

	return result, nil
}

func UpdateByUserId(userId string, userInfo UserInfo) error {
	err := collection.UpdateOne(context.Background(), bson.M{"userId": userId},
		bson.M{
			"$set": bson.M{
				"nickname":       userInfo.Nickname,
				"userStatus":     userInfo.UserStatus,
				"aboutMe":        userInfo.AboutMe,
				"education":      userInfo.Education,
				"gender":         userInfo.Gender,
				"profession":     userInfo.Profession,
				"jobExperiences": userInfo.JobExperiences,
				"fields":         userInfo.Fields,
				"mentorSkill":    userInfo.MentorSkill,
				"modifyTime":     time.Now(),
			},
		})
	if err != nil {
		return err
	}

	return nil
}

func UpdatePictureByUserId(userId string, pictureUrl string) error {
	err := collection.UpdateOne(context.Background(), bson.M{"userId": userId},
		bson.M{
			"$set": bson.M{
				"pictureUrl": pictureUrl,
				"modifyTime": time.Now(),
			},
		})
	if err != nil {
		return err
	}

	return nil
}

func UpdateAvatorUrlByUserId(userId string, avatorUrl string) error {
	err := collection.UpdateOne(context.Background(), bson.M{"userId": userId},
		bson.M{
			"$set": bson.M{
				"avatorUrl":  avatorUrl,
				"modifyTime": time.Now(),
			},
		})
	if err != nil {
		return err
	}

	return nil
}

func FindByConditionsWithPagination(page int64, fields []string, gender []string) ([]UserInfo, error) {
	batch := []UserInfo{}
	offset := page * PAGE_SIZE
	excludeStatusStage := bson.D{{Key: operator.Match, Value: bson.M{"userStatus": user_dao.USER_STUDENT}}}
	matchProfessionCategoriesStage := bson.D{{Key: operator.Match, Value: bson.M{"fields": inOrAll(fields)}}}
	matchGenderStage := bson.D{{Key: operator.Match, Value: bson.M{"gender": inOrAll(gender)}}}
	skipStage := bson.D{{Key: operator.Skip, Value: offset}}
	limitStage := bson.D{{Key: operator.Limit, Value: PAGE_SIZE}}

	err := collection.Aggregate(context.Background(), qmgo.Pipeline{
		excludeStatusStage,
		matchProfessionCategoriesStage,
		matchGenderStage,
		skipStage,
		limitStage}).All(&batch)
	if err != nil {
		return nil, err
	}

	return batch, nil
}

func FindByProfessionCategoryWithPagination(page int64, profession string) ([]UserInfo, error) {
	batch := []UserInfo{}
	offset := page * PAGE_SIZE

	err := collection.Find(context.Background(), bson.M{"profession": profession}).Skip(offset).Limit(page).All(&batch)
	if err != nil {
		return nil, err
	}

	return batch, nil
}

func FindAllWithPagination(page int64) ([]UserInfo, error) {
	batch := []UserInfo{}
	offset := page * PAGE_SIZE

	err := collection.Find(context.Background(), bson.M{}).Skip(offset).Limit(page).All(&batch)
	if err != nil {
		return nil, err
	}

	return batch, nil
}

func FindFcmAndNicknameByUserId(userId string) (dto.PushNotificationInfoDto, error) {
	result := dto.PushNotificationInfoDto{}
	err := collection.Find(context.Background(), bson.M{"userId": userId}).Select(bson.M{"fcmToken": 1, "nickname": 1}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func PushFcmByUserId(userId string, fcmToken string) error {
	err := collection.UpdateOne(context.Background(), bson.M{"userId": userId},
		bson.M{
			"$pull": bson.M{
				"fcmToken": fcmToken,
			},
		})
	if err != nil {
		return err
	}
	err = collection.UpdateOne(context.Background(), bson.M{"userId": userId},
		bson.M{
			"$push": bson.M{
				"fcmToken": fcmToken,
			},
		})
	if err != nil {
		return err
	}

	return nil
}

func PullFcmByUserId(userId string, fcmToken string) error {
	err := collection.UpdateOne(context.Background(), bson.M{"userId": userId},
		bson.M{
			"$pull": bson.M{
				"fcmToken": fcmToken,
			},
		})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func PushProfessionCategoryWithTx(ctx context.Context, professionName string, category string) error {
	_, err := collection.UpdateAll(ctx,
		bson.M{"profession": professionName},
		bson.M{"$push": bson.M{
			"professionCategories": category,
		}})
	if err != nil {
		errors.Wrap(err, "error push profession category")
		return err
	}

	return nil
}

func PullProfessionCategoryWithTx(ctx context.Context, professionName string, category string) error {
	_, err := collection.UpdateAll(ctx,
		bson.M{"profession": professionName},
		bson.M{"$pull": bson.M{
			"professionCategories": category,
		}})
	if err != nil {
		errors.Wrap(err, "error pull profession category")
		return err
	}

	return nil
}

func inOrAll(value []string) bson.M {
	if len(value) == 0 {
		return bson.M{"$exists": true}
	} else {
		return bson.M{"$in": value}
	}
}
