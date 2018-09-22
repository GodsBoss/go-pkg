package vector2d

import (
	"math"
)

// Int64 represents a 2D 64-bit integer vector.
type Int64 interface {
	X() int64
	Y() int64
}

type int64Impl struct {
	x int64
	y int64
}

// NewInt64 creates a new integer vector from a pair of coordinates.
func NewInt64(x, y int64) Int64 {
	return int64Impl{
		x: x,
		y: y,
	}
}

// RoundFloat64 rounds a floating point vector to an integer vector.
func RoundFloat64(f Float64) Int64 {
	return NewInt64(
		int64(math.Round(f.X())),
		int64(math.Round(f.Y())),
	)
}

func (i int64Impl) X() int64 {
	return i.x
}

func (i int64Impl) Y() int64 {
	return i.y
}

var zeroInt64 = int64Impl{}

// ZeroInt64 returns the integer flavor of the zero vector.
func ZeroInt64() Int64 {
	return zeroInt64
}

// AddInt64 adds together integer vectors. If no vectors are given, zero is returned.
func AddInt64(ints ...Int64) Int64 {
	value := ZeroInt64()
	for i := range ints {
		value = NewInt64(
			value.X()+ints[i].X(),
			value.Y()+ints[i].Y(),
		)
	}
	return value
}

// MultiplyInt64 multiplies an integer vector with f.
func MultiplyInt64(f int64, v Int64) Int64 {
	return int64Impl{
		x: f * v.X(),
		y: f * v.Y(),
	}
}
