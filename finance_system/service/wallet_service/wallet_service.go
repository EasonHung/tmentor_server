package wallet_service

// func CreateWallet(userId string) (error, string) {
// 	err, walletId := wallet_dao.CreateWallet(userId)
// 	if err != nil {
// 		errors.Wrap(err, "error create wallet")
// 		return err, ""
// 	}

// 	return nil, walletId
// }

// func GetWallet(userId string) (error, vo.GetWalletRes) {
// 	err, walletDao := wallet_dao.SelectByUserId(userId)
// 	if err != nil {
// 		return err, vo.GetWalletRes{}
// 	}

// 	res := vo.GetWalletRes{}
// 	res.ConvertFromWalletDao(walletDao)

// 	return nil, res
// }
