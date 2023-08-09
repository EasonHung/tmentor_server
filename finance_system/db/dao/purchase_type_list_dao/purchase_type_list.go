package purchase_list_dao

import (
	"database/sql"
	"time"
)

type PurchaseType struct {
	Id         int64
	Name       string
	SPoint     int64
	Cost       int64
	Status     string
	Desc       sql.NullString
	CreateTime time.Time
	ModifyTime time.Time
}
