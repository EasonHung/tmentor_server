package errcode

var (
	InsufficientSPointsError = NewError("F-0001", "insufficient s points error")
	UnexpectError            = NewError("9999", "unexpect error")
)
