package s_point_purchase_history_dao

// func InsertWithTx(tx *sql.Tx, entity *SPointPurchaseHistory) error {
// 	_, err := tx.Exec("INSERT INTO s_point_purchase_history_table (wallet_id, purchase_type, price, point, method, info_json, purchase_time) VALUES (?, ?, ?, ?, ?, ?, ?)",
// 		entity.WalletId, entity.PurchaseType, entity.Price, entity.Point, entity.Method, entity.InfoJson.String, entity.PurchaseTime)
// 	if err != nil {
// 		errors.Wrap(err, "error create student puchase history")
// 		return err
// 	}

// 	return nil
// }
