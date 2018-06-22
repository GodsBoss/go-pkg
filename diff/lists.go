package diff

type lists struct {
	firstLength  int
	secondLength int
	equals       func(firstIndex, secondIndex int) bool
}

// NewLists takes the length of the first list, the length of the second list and a comparison function to generate a Lists instance.
func NewLists(lengthOfFirstList, lengthOfSecondList int, equals func(firstIndex, secondIndex int) bool) Lists {
	return &lists{
		firstLength:  lengthOfFirstList,
		secondLength: lengthOfSecondList,
		equals:       equals,
	}
}

func (l *lists) Equals(firstIndex, secondIndex int) bool {
	return l.equals(firstIndex, secondIndex)
}

func (l *lists) Lengths() (int, int) {
	return l.firstLength, l.secondLength
}
