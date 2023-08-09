package user_blacklist_dao

// func InsertOne(userId string) error {
// 	_, err := connection.SQL_CLIENT.Exec(
// 		"INSERT INTO user_black_list (user_id) VALUES (?)",
// 		userId,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func DeleteOne(userId string) error {
// 	_, err := connection.SQL_CLIENT.Exec(
// 		"DELETE FROM user_black_list WHERE user_id=?",
// 		userId,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func FindAll() ([]UserBlackList, error) {
// 	batch := []UserBlackList{}

// 	rows, err := connection.SQL_CLIENT.Query("SELECT * FROM user_black_list")
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {
// 		row := UserBlackList{}

// 		if err := rows.Scan(&row.Id, &row.UserId, &row.CreateTime, &row.ModifyTime); err != nil {
// 			return nil, err
// 		}
// 		batch = append(batch, row)
// 	}

// 	return batch, nil
// }
