package finance_system_api

import (
	"encoding/json"
	"mentor/classroom/config"
	"mentor/classroom/internal_api"
	"mentor/classroom/internal_api/finance_system_api/finance_system_request"
	"mentor/classroom/internal_api/finance_system_api/finance_system_response"

	"github.com/pkg/errors"
)

var financeApiHandler internal_api.InternalApiHandler

func init() {
	financeApiHandler = internal_api.NewApiHandler(config.GLOBAL_CONFIG.FinanceSystem.BaseUrl)
}

func PayClassBill(classId string, costSPoints int, mentorId string, studentId string) (error, finance_system_response.PayClassBillRes) {
	payClassBillResp := finance_system_response.PayClassBillRes{}

	reqBody := finance_system_request.NewPayClassBillReq(classId, costSPoints, mentorId, studentId)
	resBody, err := financeApiHandler.Post("/purchase/class/pay", reqBody)
	if err != nil {
		return err, payClassBillResp
	}

	err = json.Unmarshal(resBody, &payClassBillResp)
	if err != nil {
		return errors.WithStack(err), payClassBillResp
	}

	return nil, payClassBillResp
}
