package wallet_dao

import (
	"database/sql"
	"time"
)

type Wallet struct {
	WalletId     string
	UserId       string
	StudentPoint int64
	TeacherPoint int64
	InfoJson     sql.NullString
	CreateTime   time.Time
	ModifyTime   time.Time
}
