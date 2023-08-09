package classroom_api

type GetStudentCountRes struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    GetStudentCountResData `json:"data"`
}

type GetStudentCountResData struct {
	UserId string `json:"userId"`
	Count  int    `json:"count"`
}
