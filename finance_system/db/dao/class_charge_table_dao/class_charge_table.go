package class_charge_table_dao

import (
	"database/sql"
	"time"
)

type ClassChargeTable struct {
	Id           int64
	ClassId      string
	SWalletId    string
	TWalletId    string
	SPoint       int64
	TPoint       int64
	PurchaseTime time.Time
	InfoJson     sql.NullString
	CreateTime   time.Time
	ModifyTime   time.Time
}
