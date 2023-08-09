package user_dao

import (
	"database/sql"
	"time"
)

// user status enum
const (
	USER_INIT    = "initial"
	USER_STUDENT = "student"
)

type User struct {
	UserId       string
	UserStatus   string
	LoginInfoId  sql.NullInt64
	ThirdPartyId sql.NullString
	WalletId     string
	CreateTime   time.Time
	ModifyTime   time.Time
}
