package interval

// Int64 represents an integer interval. The zero value is an interval without bounds,
// i.e. from negative infinity to positive infinity.
type Int64 struct {
	left  *leftInt64
	right *rightInt64
}

type leftInt64 struct {
	left int64
	BoundsType
}

func (left leftInt64) includes(i int64) bool {
	return i > left.left || (left.includeBound() && i == left.left)
}

type rightInt64 struct {
	right int64
	BoundsType
}

func (right rightInt64) includes(i int64) bool {
	return i < right.right || (right.includeBound() && i == right.right)
}

func (interval Int64) Inside(i int64) bool {
	return (interval.left == nil || interval.left.includes(i)) && (interval.right == nil || interval.right.includes(i))
}

func (interval Int64) Left(left int64, boundsType BoundsType) Int64 {
	interval.left = &leftInt64{
		left:       left,
		BoundsType: boundsType,
	}
	return interval
}

func (interval Int64) Right(right int64, boundsType BoundsType) Int64 {
	interval.right = &rightInt64{
		right:      right,
		BoundsType: boundsType,
	}
	return interval
}

type BoundsType interface {
	includeBound() bool
}

type boundsType bool

func (t boundsType) includeBound() bool {
	return bool(t)
}

const includeBound boundsType = true
const excludeBound boundsType = false

func Inclusive() BoundsType {
	return includeBound
}

func Exclusive() BoundsType {
	return excludeBound
}
