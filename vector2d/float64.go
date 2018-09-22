package vector2d

// Float64 is a 2D 64-bit floating-point vector.
type Float64 interface {
	X() float64
	Y() float64
}

type float64Impl struct {
	x float64
	y float64
}

// NewFloat64 creates a new floating-point vector from a pair of coordinates.
func NewFloat64(x, y float64) Float64 {
	return float64Impl{
		x: x,
		y: y,
	}
}

// Float64FromInt64 converts an integer vector to a floating-point vector.
func Float64FromInt64(i Float64) Float64 {
	return NewFloat64(float64(i.X()), float64(i.Y()))
}

func (f float64Impl) X() float64 {
	return f.x
}

func (f float64Impl) Y() float64 {
	return f.y
}

var zeroFloat64 = float64Impl{}

// ZeroFloat64 returns the floating-point flavor of the zero vector.
func ZeroFloat64() Float64 {
	return zeroFloat64
}

// AddFloat64s adds together floating-point vectors. If no vectors are given,
// zero is returned.
func AddFloat64s(floats ...Float64) Float64 {
	value := ZeroFloat64()
	for i := range floats {
		value = NewFloat64(
			value.X()+floats[i].X(),
			value.Y()+floats[i].Y(),
		)
	}
	return value
}

// ScaleFloat64 scales a floating-point vector by f.
func ScaleFloat64(f float64, fl Float64) Float64 {
	return NewFloat64(f*fl.X(), f*fl.Y())
}
