package s_point_purchase_history_dao

import (
	"database/sql"
	"time"
)

type SPointPurchaseHistory struct {
	WalletId     string
	PurchaseType string
	Price        int64
	Point        int64
	Method       string
	InfoJson     sql.NullString
	PurchaseTime time.Time
	CreateTime   time.Time
	ModifyTime   time.Time
}
