package internal_error

import "fmt"

type OutOfLimitError struct {
}

func (e OutOfLimitError) Error() string {
	return fmt.Sprintf("exeed the limit")
}
