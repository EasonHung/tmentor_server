package login_info_dao

import "time"

type LoginInfo struct {
	Id         int64
	Account    string
	Password   string
	CreateTime time.Time
	ModifyTime time.Time
}
