package zero_test

import (
	"github.com/GodsBoss/go-pkg/zero"

	"fmt"
)

func ExampleGuard() {
	// Proper initialization, no problem.
	NewUselessZeroValue().Greet()

	func() {
		defer func() {
			fmt.Println(recover())
		}()

		// Zero value will panic.
		(&UselessZeroValue{}).Greet()
	}()

	// Output:
	// Hello, world!
	// UselessZeroValue is uninitialized
}

// UselessZeroValue must be initialized with NewUselessZeroValue. The zero value is not useful!
type UselessZeroValue struct {
	guard *zero.Guard
}

// NewUselessZeroValue creates a usable instance of UselessZeroValue.
func NewUselessZeroValue() *UselessZeroValue {
	return &UselessZeroValue{
		guard: &zero.Guard{},
	}
}

// Greet greets the world.
func (value *UselessZeroValue) Greet() {
	value.guard.Check("UselessZeroValue")
	fmt.Println("Hello, world!")
}
