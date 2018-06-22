package diff

// Complex128s diffs two lists of complex128s.
func Complex128s(first []complex128, second []complex128) *Result {
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

// Complex64s diffs two lists of complex64s.
func Complex64s(first []complex64, second []complex64) *Result {
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

// Float32s diffs two lists of float64s.
func Float32s(first []float32, second []float32) *Result {
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

// Float64s diffs two lists of float64s.
func Float64s(first []float64, second []float64) *Result {
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

// Runes diffs two lists of runes.
func Runes(first []rune, second []rune) *Result {
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

// Strings diffs two lists of strings.
func Strings(first []string, second []string) *Result {
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

// Uints diffs two lists of uints.
func Uints(first []uint, second []uint) *Result {
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
