package internal_error

import "fmt"

type InsufficientSPointsError struct {
}

func (e InsufficientSPointsError) Error() string {
	return fmt.Sprintf("insufficient spoints error")
}
