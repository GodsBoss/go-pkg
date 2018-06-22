package diff

// Lists wraps two lists. It is consumed by Diff().
type Lists interface {
	// Equals returns wether the items with the given indexes in the first and second list are equal.
	Equals(firstIndex, secondIndex int) bool

	// Lengths returns the length of the first and second list, in that order.
	Lengths() (int, int)
}

// Result is the result of diffing two lists.
type Result struct {
	inFirstListButNotSecond []int
	inSecondListButNotFirst []int
	inBothLists             []*InBothListsIndex
}

// InFirstListButNotSecond returns the indexes of values found in the first list, but not the second, in ascending order, i.e. lower indexes first.
func (result *Result) InFirstListButNotSecond() []int {
	return result.inFirstListButNotSecond
}

// InSecondListButNotFirst returns the indexes of values found in the second list, but not the first, in ascending order, i.e. lower indexes first.
func (result *Result) InSecondListButNotFirst() []int {
	return result.inSecondListButNotFirst
}

// InBothLists returns the indexes of values found in both lists, in arbitrary order.
func (result *Result) InBothLists() []*InBothListsIndex {
	return result.inBothLists
}

// InBothListsIndex contains two indexes, one for the first list, one for the second. The values for these indexes are equal.
type InBothListsIndex struct {
	first  int
	second int
}

// First returns the index for the first list.
func (index *InBothListsIndex) First() int {
	return index.first
}

// Second returns the index for the second list.
func (index *InBothListsIndex) Second() int {
	return index.second
}

// Diff diffs lists and returns the result.
func Diff(lists Lists) *Result {
	result := &Result{
		inFirstListButNotSecond: []int{},
		inSecondListButNotFirst: []int{},
		inBothLists:             []*InBothListsIndex{},
	}
	firstLength, secondLength := lists.Lengths()
	for firstIndex := 0; firstIndex < firstLength; firstIndex++ {
		found := false
		for secondIndex := 0; secondIndex < secondLength; secondIndex++ {
			if lists.Equals(firstIndex, secondIndex) {
				found = true
				result.inBothLists = append(
					result.inBothLists,
					&InBothListsIndex{
						first:  firstIndex,
						second: secondIndex,
					},
				)
				break
			}
		}
		if !found {
			result.inFirstListButNotSecond = append(result.inFirstListButNotSecond, firstIndex)
		}
	}
	for secondIndex := 0; secondIndex < secondLength; secondIndex++ {
		found := false
		for firstIndex := 0; firstIndex < firstLength; firstIndex++ {
			if lists.Equals(firstIndex, secondIndex) {
				found = true
				break
			}
		}
		if !found {
			result.inSecondListButNotFirst = append(result.inSecondListButNotFirst, secondIndex)
		}
	}
	return result
}
