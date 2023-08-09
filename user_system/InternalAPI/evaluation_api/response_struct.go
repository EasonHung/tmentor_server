package evaluation_api

type GetScoreAndCountRes struct {
	Code    string                  `json:"code"`
	Message string                  `json:"message"`
	Data    GetScoreAndCountResData `json:"data"`
}

type GetScoreAndCountResData struct {
	UserId string  `json:"userId"`
	Score  float32 `json:"score"`
	Count  int     `json:"count"`
}
