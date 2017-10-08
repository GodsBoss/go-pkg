package copyguard_test

import (
	"github.com/GodsBoss/go-pkg/copyguard"

	"fmt"
)

type MustNotBeCopied struct {
	guard copyguard.Guard
}

func (m *MustNotBeCopied) Show(who string) {
	m.guard.Check(m, fmt.Sprintf("%s used after copy!"))
	fmt.Printf("This is %s!\n", who)
}

func ExampleGuard_Check_noPanicOnUseAfterCopy() {
	original := MustNotBeCopied{}
	clone := original

	original.Show("the original")
	clone.Show("the clone")

	// Output:
	// This is the original!
	// This is the clone!
}
