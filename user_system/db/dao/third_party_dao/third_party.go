package third_party_dao

import (
	"database/sql"
	"time"
)

type ThirdParty struct {
	ThirdPartyId   sql.NullString
	ThirdPartyInfo string
	CreateTime     time.Time
	ModifyTime     time.Time
}
