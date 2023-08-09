package dto

type PushNotificationInfoDto struct {
	FcmToken []string `bson:"fcmToken"`
	Nickname string   `bson:"nickname"`
}
