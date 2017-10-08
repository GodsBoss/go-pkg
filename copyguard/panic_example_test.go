package copyguard_test

import (
	"github.com/GodsBoss/go-pkg/copyguard"

	"fmt"
)

type ShouldNotBeCopied struct {
	guard copyguard.Guard
}

func (safe *ShouldNotBeCopied) DoSomething() {
	safe.guard.Check(safe, "Non-copyable struct")
}

func ExampleGuard_Check_panicOnCopyAfterUse() {
	original := ShouldNotBeCopied{}
	original.DoSomething()
	clone := original

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		clone.DoSomething()
	}()

	// Output:
	// Non-copyable struct
}
