package errcode

var (
	UserNotFoundError = NewError("CS-0001", "user not found")
	OutOfLimitError   = NewError("CS-002", "out of limit")
	DuplicateKeyError = NewError("CS-003", "duplicate key")
	DataNotFindError = NewError("CS-004", "data not find")
	WrongStatusError = NewError("CS-005", "wrong status")
	UnexpectError     = NewError("9999", "unexpect error")
)
