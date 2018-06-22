package diff

// Ints diffs two lists of ints.
func Ints(first []int, second []int) *Result {
	return Diff(
		NewLists(
			len(first),
			len(second),
			func(firstIndex, secondIndex int) bool {
				return first[firstIndex] == second[secondIndex]
			},
		),
	)
}
