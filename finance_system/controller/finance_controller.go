package controller

import (
	"mentor_app/finance_system/vo"

	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	walletRoute := route.Group("/wallet")
	{
		walletRoute.POST("/new", CreateWallet) //內部api
		walletRoute.GET("/", getUserWallet)
	}

	purchaseRoute := route.Group("/purchase")
	{
		purchaseRoute.POST("/sPoint", PurchaseSPoint)
		purchaseRoute.POST("/class/pay", PayClassBill) //內部api
	}
}

func getUserWallet(c *gin.Context) {

	c.JSON(200, nil)
	return
}

func PayClassBill(c *gin.Context) {
	var req vo.PayClassBillReq
	c.BindJSON(&req)

	c.JSON(200, vo.NewSuccessResponse(nil))
	return
}

func CreateWallet(c *gin.Context) {
	var req vo.CreateWalletReq
	c.BindJSON(&req)
	c.JSON(200, nil)
}

func PurchaseSPoint(c *gin.Context) {
	var req vo.PurchaseSPointReq
	c.BindJSON(&req)

	c.JSON(200, "OK")
}
