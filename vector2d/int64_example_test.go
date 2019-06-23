package vector2d_test

import (
	"fmt"

	"github.com/GodsBoss/go-pkg/vector2d"
)

func ExampleAddInt64() {
	v := vector2d.AddInt64(vector2d.NewInt64(2, -2), vector2d.NewInt64(3, 1), vector2d.NewInt64(-1, 8))
	fmt.Printf("Vector is now (%d, %d).\n", v.X(), v.Y())

	// Output:
	// Vector is now (4, 7).
}

func ExampleMultiplyInt64() {
	v := vector2d.MultiplyInt64(6, vector2d.NewInt64(3, -5))
	fmt.Printf("Vector is now (%d, %d).\n", v.X(), v.Y())

	// Output:
	// Vector is now (18, -30).
}
