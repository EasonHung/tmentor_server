package user_dao

import (
	"mentor_app/user_system/db/connection"

	"github.com/qiniu/qmgo"
)

var collection *qmgo.Collection

func init() {
	collection = connection.MONGO_CLIENT.Database("user_info_system").Collection("user_info")
}

// func CreateWithTx(tx *sql.Tx, entity *User) {
// 	_, err := tx.Exec("INSERT INTO user_tab (user_id, user_status, login_info_id, third_party_id, wallet_id) VALUES (?, ?, ?, ?, ?)",
// 		entity.UserId, entity.UserStatus, entity.LoginInfoId.Int64, entity.ThirdPartyId.String, entity.WalletId)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func GetByThirdPartyId(thirdPartyId string) (User, error) {
// 	var result User
// 	row := connection.SQL_CLIENT.QueryRow("SELECT * FROM user_tab WHERE third_party_id = ?", thirdPartyId)
// 	err := row.Scan(&result.UserId, &result.UserStatus, &result.LoginInfoId, &result.ThirdPartyId, &result.WalletId, &result.CreateTime, &result.ModifyTime)
// 	if err != nil {
// 		return result, err
// 	}

// 	return result, nil
// }

// func GetByLoginInfoId(loginInfoId int64) (User, error) {
// 	var result User
// 	row := connection.SQL_CLIENT.QueryRow("SELECT * FROM user_tab WHERE login_info_id = ?", loginInfoId)
// 	err := row.Scan(&result.UserId, &result.UserStatus, &result.LoginInfoId, &result.ThirdPartyId, &result.WalletId, &result.CreateTime, &result.ModifyTime)
// 	if err != nil {
// 		return result, err
// 	}

// 	return result, nil
// }

// func GetByUserId(userId string) (User, error) {
// 	var result User
// 	row := connection.SQL_CLIENT.QueryRow("SELECT * FROM user_tab WHERE user_id = ?", userId)
// 	err := row.Scan(&result.UserId, &result.UserStatus, &result.LoginInfoId, &result.ThirdPartyId, &result.WalletId, &result.CreateTime, &result.ModifyTime)
// 	if err != nil {
// 		return result, err
// 	}

// 	return result, nil
// }

// func UpdateStatusByUserId(userId string, status string) error {
// 	_, err := connection.SQL_CLIENT.Exec(
// 		"UPDATE user_tab SET user_status = ? WHERE user_id = ?",
// 		status, userId,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
