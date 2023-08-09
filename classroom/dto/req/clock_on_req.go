package req

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type ClockOnReq struct {
	ClassId   string `json:"classId"`
}

func JsonStrToClockOnReq(jsonStr string) (error, ClockOnReq) {
	res := ClockOnReq{}
	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		return errors.WithStack(err), res
	}
	return nil, res
}
