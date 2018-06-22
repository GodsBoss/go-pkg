package zero

import (
	"fmt"
)

// Guard simplifies aborting when structs with unusable zero values are not initialized properly.
type Guard struct{}

// Check panics with an appropriate message when guard is nil.
func (guard *Guard) Check(name string) {
	if guard == nil {
		panic(fmt.Sprintf(MessagePattern, name))
	}
}

// MessagePattern is the panic message with a placeholder for the name.
const MessagePattern = "%s is uninitialized"
