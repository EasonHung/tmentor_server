package vo

type GetScoreCountRes struct {
	UserId string  `json:"userId"`
	Score  float64 `json:"score"`
	Count  int     `json:"count"`
}
