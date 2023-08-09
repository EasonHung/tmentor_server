package DateUtils

import "time"

func ShiftMinutes(minutes int) time.Time {
	return time.Now().Local().Add(time.Minute * time.Duration(minutes))
}

func ShiftDays(days int) time.Time {
	return time.Now().Local().AddDate(0, 0, days)
}
