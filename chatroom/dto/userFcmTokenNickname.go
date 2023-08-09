package dto

type UserFcmTokenNickname struct {
	FcmToken string `bson:"fcmToken"`
	Nickname string `bson:"nickname"`
}
