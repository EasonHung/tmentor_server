package vo

import "mentor_app/finance_system/db/dao/wallet_dao"

type GetWalletRes struct {
	WalletId     string `json:"walletId"`
	StudentPoint int64  `json:"studentPoint"`
	TeacherPoint int64  `json:"teacherPoint"`
	InfoJson     string `json:"infoJson"`
}

func (this *GetWalletRes) ConvertFromWalletDao(walletDao wallet_dao.Wallet) {
	this.WalletId = walletDao.WalletId
	this.StudentPoint = walletDao.StudentPoint
	this.TeacherPoint = walletDao.TeacherPoint
	this.InfoJson = walletDao.InfoJson.String
}
