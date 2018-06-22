package diff

// Ints diffs two lists of ints.
func Ints(first []int, second []int) *Result {
	return Diff(
		&intLists{
			first:  first,
			second: second,
		},
	)
}

type intLists struct {
	first  []int
	second []int
}

func (lists *intLists) Equals(firstIndex, secondIndex int) bool {
	return lists.first[firstIndex] == lists.second[secondIndex]
}

func (lists *intLists) Lengths() (int, int) {
	return len(lists.first), len(lists.second)
}
