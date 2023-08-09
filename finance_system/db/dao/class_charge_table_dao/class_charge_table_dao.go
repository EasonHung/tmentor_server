package class_charge_table_dao

// func InsertWithTx(tx *sql.Tx, entity *ClassChargeTable) error {
// 	_, err := tx.Exec("INSERT INTO class_charge_table (class_id, s_wallet_id, t_wallet_id, s_point, t_point, purchase_time, info_json) VALUES (?, ?, ?, ?, ?, ?, ?)",
// 		entity.ClassId, entity.SWalletId, entity.TWalletId, entity.SPoint, entity.TPoint, entity.PurchaseTime, entity.InfoJson)
// 	if err != nil {
// 		errors.Wrap(err, "error create s to t puchase history")
// 		return err
// 	}

// 	return nil
// }
