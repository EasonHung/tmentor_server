package exchange_rate_dao

// func GetExchangeRateByName(name string) (error, float64) {
// 	var result ExchangeRate
// 	row := connection.SQL_CLIENT.QueryRow("SELECT * FROM exchange_rate_table WHERE name = ?", name)
// 	err := row.Scan(&result.Id, &result.Name, &result.ExchangeRate)
// 	if err != nil {
// 		err = errors.Wrap(err, "error get Exchange rate")
// 		return err, 0
// 	}

// 	return nil, result.ExchangeRate
// }
