package wallet_dao

// func CreateWallet(userId string) (error, string) {
// 	walletId := xid.New().String()
// 	_, err := connection.SQL_CLIENT.Exec(
// 		"INSERT INTO wallet_table (wallet_id, user_id) VALUES (?, ?)",
// 		walletId,
// 		userId,
// 	)
// 	if err != nil {
// 		err = errors.Wrap(err, "error create wallet")
// 		return err, ""
// 	}

// 	return nil, walletId
// }

// func IncreaseSPointByTxWithWalletId(tx *sql.Tx, walletId string, sPoint string) error {
// 	// 這句會有 injection 的疑慮 不能給外部打
// 	_, err := tx.Exec("UPDATE wallet_table SET student_point = student_point + "+sPoint+" WHERE wallet_id = ?", walletId)
// 	if err != nil {
// 		err = errors.Wrap(err, "error update sPoint")
// 		return err
// 	}

// 	return nil
// }

// func SelectByWalletId(walletId string) (error, Wallet) {
// 	var result Wallet
// 	row := connection.SQL_CLIENT.QueryRow("SELECT * FROM wallet_table WHERE wallet_id = ?", walletId)
// 	err := row.Scan(&result.WalletId, &result.UserId, &result.StudentPoint, &result.TeacherPoint, &result.InfoJson, &result.CreateTime, &result.ModifyTime)
// 	if err != nil {
// 		err = errors.Wrap(err, "error get Exchange rate")
// 		return err, result
// 	}

// 	return nil, result
// }

// func SelectByUserId(userId string) (error, Wallet) {
// 	var result Wallet
// 	row := connection.SQL_CLIENT.QueryRow("SELECT * FROM wallet_table WHERE user_id = ?", userId)
// 	err := row.Scan(&result.WalletId, &result.UserId, &result.StudentPoint, &result.TeacherPoint, &result.InfoJson, &result.CreateTime, &result.ModifyTime)
// 	if err != nil {
// 		err = errors.Wrap(err, "error get Exchange rate")
// 		return err, result
// 	}

// 	return nil, result
// }

// func DecreaseSPointByTxWithWalletId(tx *sql.Tx, walletId string, sPoint string) error {
// 	// 這句會有 injection 的疑慮 不能給外部打
// 	_, err := tx.Exec("UPDATE wallet_table SET student_point = student_point - "+sPoint+" WHERE wallet_id = ?", walletId)
// 	if err != nil {
// 		err = errors.Wrap(err, "error update sPoint")
// 		return err
// 	}

// 	return nil
// }

// func IncreaseTPointByTxWithWalletId(tx *sql.Tx, walletId string, tPoint string) error {
// 	// 這句會有 injection 的疑慮 不能給外部打
// 	_, err := tx.Exec("UPDATE wallet_table SET teacher_point = teacher_point + "+tPoint+" WHERE wallet_id = ?", walletId)
// 	if err != nil {
// 		err = errors.Wrap(err, "error update sPoint")
// 		return err
// 	}

// 	return nil
// }
