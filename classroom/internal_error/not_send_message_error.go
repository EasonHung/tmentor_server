package internal_error

import "fmt"

type NotSendMessageError struct {
}

func (e NotSendMessageError) Error() string {
	return fmt.Sprintf("not send message error")
}
