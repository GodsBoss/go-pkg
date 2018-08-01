package interval

// Int64 represents an integer interval. The zero value is an interval without bounds,
// i.e. from negative infinity to positive infinity.
type Int64 struct {
	lower *lowerInt64
	upper *upperInt64
}

type lowerInt64 struct {
	value int64
	BoundsType
}

func (lower lowerInt64) includes(i int64) bool {
	return i > lower.value || (lower.includeBound() && i == lower.value)
}

type upperInt64 struct {
	value int64
	BoundsType
}

func (upper upperInt64) includes(i int64) bool {
	return i < upper.value || (upper.includeBound() && i == upper.value)
}

func (interval Int64) Inside(i int64) bool {
	return (interval.lower == nil || interval.lower.includes(i)) && (interval.upper == nil || interval.upper.includes(i))
}

func (interval Int64) Lower(bound int64, boundsType BoundsType) Int64 {
	interval.lower = &lowerInt64{
		value:      bound,
		BoundsType: boundsType,
	}
	return interval
}

func (interval Int64) Upper(bound int64, boundsType BoundsType) Int64 {
	interval.upper = &upperInt64{
		value:      bound,
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
