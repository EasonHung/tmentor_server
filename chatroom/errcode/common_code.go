package errcode

var (
	DuplicateAddConversation = NewError("CR-0001", "duplicate adding conversation")
	UnexpectError            = NewError("9999", "unexpect error")
)
