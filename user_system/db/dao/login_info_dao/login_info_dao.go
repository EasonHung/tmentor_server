package login_info_dao

// func CreateWithTx(tx *sql.Tx, entity *LoginInfo) (int64, error) {
// 	res, err := tx.Exec("INSERT INTO login_info_table (account, password) VALUES (?, ?)",
// 		entity.Account, entity.Password)
// 	if err != nil {
// 		errors.Wrap(err, "error insert user")
// 		return 0, err
// 	}

// 	resId, err := res.LastInsertId()
// 	if err != nil {
// 		errors.Wrap(err, "error insert user")
// 		return resId, err
// 	}

// 	return resId, nil
// }

// func GetByAccount(account string) (LoginInfo, error) {
// 	var result LoginInfo
// 	row := connection.SQL_CLIENT.QueryRow("SELECT * FROM login_info_table WHERE account = ?", account)
// 	err := row.Scan(&result.Id, &result.Account, &result.Password, &result.CreateTime, &result.ModifyTime)
// 	if err != nil {
// 		errors.Wrap(err, "error get user by account")
// 		return result, err
// 	}

// 	return result, nil
// }
