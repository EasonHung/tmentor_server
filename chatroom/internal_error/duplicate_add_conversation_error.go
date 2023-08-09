package internal_error

import "fmt"

type DuplicateAddConversationError struct {
	ConversationId string
}

func (e DuplicateAddConversationError) Error() string {
	return fmt.Sprintf("duplicate adding conversation")
}
