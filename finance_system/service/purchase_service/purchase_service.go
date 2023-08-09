package purchase_service

// func PayClassBill(classId string, teacherUserId string, studentUserId string, costSPoint int) error {
// 	err, teacherWalletId := user_api.GetWalletId(teacherUserId)
// 	if err != nil {
// 		err = errors.Wrap(err, "error get teacher WalletId")
// 		return err
// 	}

// 	err, studentWalletId := user_api.GetWalletId(studentUserId)
// 	if err != nil {
// 		err = errors.Wrap(err, "error get student WalletId")
// 		return err
// 	}

// 	err, studentWallet := wallet_dao.SelectByWalletId(studentWalletId)
// 	if err != nil {
// 		err = errors.Wrap(err, "error get student wallet")
// 		return err
// 	}
// 	if studentWallet.StudentPoint < int64(costSPoint) {
// 		return errors.Wrap(internal_error.InsufficientSPointsError{}, "insufficient s points")
// 	}

// 	err, earnedTPoint := exchangeTPoint(costSPoint)
// 	if err != nil {
// 		err = errors.Wrap(err, "error get exchanged earnedTPoint")
// 		return err
// 	}

// 	tx, err := connection.SQL_CLIENT.Begin()
// 	defer tx.Rollback()
// 	if err != nil {
// 		panic(err)
// 	}

// 	purchaseHistory := makeSToTPurchaseHistoryEntity(classId, studentWalletId, teacherWalletId, int64(costSPoint), int64(earnedTPoint))
// 	err = s_to_t_purchase_history_dao.InsertWithTx(tx, &purchaseHistory)
// 	if err != nil {
// 		err = errors.Wrap(err, "error insert purchase history")
// 		return err
// 	}

// 	err = wallet_dao.DecreaseSPointByTxWithWalletId(tx, studentWalletId, strconv.FormatInt(int64(costSPoint), 10))
// 	if err != nil {
// 		err = errors.Wrap(err, "error update wallet")
// 		return err
// 	}

// 	err = wallet_dao.IncreaseTPointByTxWithWalletId(tx, teacherWalletId, strconv.FormatInt(int64(earnedTPoint), 10))
// 	if err != nil {
// 		err = errors.Wrap(err, "error update wallet")
// 		return err
// 	}

// 	tx.Commit()
// 	return nil
// }

// func BuyStudentPoint(userToken string, method string, purchaseType string, paymentInfo string) error {
// 	err, userMap := user_api.VerifyToken(userToken)
// 	if err != nil {
// 		errors.Wrap(err, "error verify userToken")
// 		return err
// 	}

// 	log.Logger.Info("walletId: " + userMap["walletId"].(string))
// 	err, purchaseTypeInfo := purchase_list_dao.GetByName(purchaseType)
// 	if err != nil {
// 		err = errors.Wrap(err, "error get purchase type info")
// 		return err
// 	}

// 	tx, err := connection.SQL_CLIENT.Begin()
// 	defer tx.Rollback()
// 	if err != nil {
// 		panic(err)
// 	}

// 	purchaseHistory := makePurchaseHistoryEntity(userMap["walletId"].(string), method, paymentInfo, purchaseTypeInfo)
// 	err = s_point_purchase_history_dao.InsertWithTx(tx, &purchaseHistory)
// 	if err != nil {
// 		err = errors.Wrap(err, "error insert s point purchase history")
// 		return err
// 	}

// 	err = wallet_dao.IncreaseSPointByTxWithWalletId(tx, userMap["walletId"].(string), strconv.FormatInt(purchaseTypeInfo.SPoint, 10))
// 	if err != nil {
// 		err = errors.Wrap(err, "error update wallet")
// 		return err
// 	}

// 	tx.Commit()

// 	return nil
// }

// func makePurchaseHistoryEntity(walletId string, method string, paymentInfo string, purchaseTypeInfo purchase_list_dao.PurchaseType) s_point_purchase_history_dao.SPointPurchaseHistory {
// 	entity := s_point_purchase_history_dao.SPointPurchaseHistory{
// 		WalletId:     walletId,
// 		PurchaseType: purchaseTypeInfo.Name,
// 		Price:        purchaseTypeInfo.Cost,
// 		Point:        purchaseTypeInfo.SPoint,
// 		Method:       method,
// 		InfoJson:     sql.NullString{String: paymentInfo},
// 		PurchaseTime: time.Now(),
// 	}

// 	return entity
// }

// func makeSToTPurchaseHistoryEntity(classId string, sWalletId string, tWalletId string, sPoint int64, tPoint int64) class_charge_table_dao.ClassChargeTable {
// 	entity := class_charge_table_dao.ClassChargeTable{
// 		ClassId:      classId,
// 		SWalletId:    sWalletId,
// 		TWalletId:    tWalletId,
// 		SPoint:       sPoint,
// 		TPoint:       tPoint,
// 		InfoJson:     sql.NullString{String: ""},
// 		PurchaseTime: time.Now(),
// 	}

// 	return entity
// }

// func exchangeTPoint(sPoint int) (error, int) {
// 	floatSPoint := float64(sPoint)
// 	err, exchangeRate := exchange_rate_dao.GetExchangeRateByName("s_to_t")
// 	if err != nil {
// 		err = errors.Wrap(err, "error get exchange rate")
// 		return err, 0
// 	}

// 	floatTPoint := floatSPoint * exchangeRate
// 	return nil, int(math.Round(floatTPoint))
// }
