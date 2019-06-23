package vector2d_test

import (
	"fmt"

	"github.com/GodsBoss/go-pkg/vector2d"
)

func ExampleAddFloat64s() {
	v := vector2d.AddFloat64s(vector2d.NewFloat64(1.5, -2.5), vector2d.NewFloat64(-0.5, 1.0))
	fmt.Printf("Vector is now (%.1f, %.1f).\n", v.X(), v.Y())

	// Output:
	// Vector is now (1.0, -1.5).
}
