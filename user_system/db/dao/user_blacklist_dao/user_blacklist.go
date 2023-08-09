package user_blacklist_dao

import "time"

type UserBlackList struct {
	Id         int64
	UserId     string
	CreateTime time.Time
	ModifyTime time.Time
}
