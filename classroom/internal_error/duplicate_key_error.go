package internal_error

import "fmt"

type DuplicateKeyError struct {
}

func (e DuplicateKeyError) Error() string {
	return fmt.Sprintf("duplicate key")
}
