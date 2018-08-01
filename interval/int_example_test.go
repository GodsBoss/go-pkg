package interval_test

import (
	"fmt"
	"math"

	"github.com/GodsBoss/go-pkg/interval"
)

func ExampleInt64_noBounds() {
	inter := interval.Int64{}

	values := []int64{
		math.MinInt64,
		0,
		math.MaxInt64,
	}

	for i := range values {
		if inter.Inside(values[i]) {
			fmt.Printf("%d is inside\n", values[i])
		}
	}

	// Output:
	// -9223372036854775808 is inside
	// 0 is inside
	// 9223372036854775807 is inside
}

func ExampleInt64_leftBound() {
	base := interval.Int64{}
	incl := base.Left(100, interval.Inclusive())
	excl := base.Left(100, interval.Exclusive())

	values := []int64{
		99,
		100,
		1000000000,
	}

	for i := range values {
		if incl.Inside(values[i]) {
			fmt.Printf("%d is inside (inclusive left bound)\n", values[i])
		}
		if excl.Inside(values[i]) {
			fmt.Printf("%d is inside (exclusive left bound)\n", values[i])
		}
	}

	// Output:
	// 100 is inside (inclusive left bound)
	// 1000000000 is inside (inclusive left bound)
	// 1000000000 is inside (exclusive left bound)
}

func ExampleInt64_rightBound() {
	base := interval.Int64{}
	incl := base.Right(250, interval.Inclusive())
	excl := base.Right(250, interval.Exclusive())

	values := []int64{
		-50000,
		250,
		251,
	}

	for i := range values {
		if incl.Inside(values[i]) {
			fmt.Printf("%d is inside (inclusive right bound)\n", values[i])
		}
		if excl.Inside(values[i]) {
			fmt.Printf("%d is inside (exclusive right bound)\n", values[i])
		}
	}

	// Output:
	// -50000 is inside (inclusive right bound)
	// -50000 is inside (exclusive right bound)
	// 250 is inside (inclusive right bound)
}
