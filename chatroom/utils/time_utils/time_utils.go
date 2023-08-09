package time_utils

import "time"

func TimeInTaipei(t time.Time) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Taipei")
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}
