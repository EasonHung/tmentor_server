package purchase_list_dao

// func GetByName(name string) (error, PurchaseType) {
// 	var result PurchaseType
// 	row := connection.SQL_CLIENT.QueryRow("SELECT * FROM purchase_type_list WHERE name = ?", name)
// 	err := row.Scan(&result.Id, &result.Name, &result.SPoint, &result.Cost, &result.Desc, &result.Status, &result.CreateTime, &result.ModifyTime)
// 	if err != nil {
// 		err = errors.Wrap(err, "error get Purchase type")
// 		return err, result
// 	}

// 	return nil, result
// }
