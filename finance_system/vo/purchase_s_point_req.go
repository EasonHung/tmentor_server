package vo

type PurchaseSPointReq struct {
	Method       string `json:"method"`
	PaymentInfo  string `json:"paymentInfo"`
	PurchaseType string `json:"purchaseType"`
}
