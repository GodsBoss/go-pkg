package diff_test

import (
	"github.com/GodsBoss/go-pkg/diff"

	"fmt"
)

func ExampleInts() {
	first := []int{2, 3, 11, 7, 5}
	second := []int{4, 5, 6, 7, 8, 12}

	result := diff.Ints(first, second)

	fmt.Println("Values found in the first list, but not the second:")
	for _, index := range result.InFirstListButNotSecond() {
		fmt.Printf("- %d\n", first[index])
	}
	fmt.Println("Values found in the second list, but not the first:")
	for _, index := range result.InSecondListButNotFirst() {
		fmt.Printf("- %d\n", second[index])
	}

	// Check values found in both lists.
	for _, index := range result.InBothLists() {
		if first[index.First()] != second[index.Second()] {
			fmt.Printf("Expected equal values, but got %d and %d.\n", first[index.First()], second[index.Second()])
		}
	}

	// Output:
	// Values found in the first list, but not the second:
	// - 2
	// - 3
	// - 11
	// Values found in the second list, but not the first:
	// - 4
	// - 6
	// - 8
	// - 12
}
