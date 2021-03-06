package affinematrix2d_test

import (
	"github.com/GodsBoss/go-pkg/affinematrix2d"

	"testing"
)

func newVector(x, y float64) testVector {
	return testVector{
		x: x,
		y: y,
	}
}

func sameTolerance(tolerance float64) testVector {
	return newVector(tolerance, tolerance)
}

type testVector struct {
	x float64
	y float64
}

func assertFloat64Within(t *testing.T, name string, expected, tolerance, actual float64) {
	if actual < expected-tolerance || actual > expected+tolerance {
		t.Errorf("expected %s to be %f +- %f, but got %f", name, expected, tolerance, actual)
	}
}

func assertVectorWithin(t *testing.T, name string, expected, tolerance testVector, actual affinematrix2d.Vector) {
	assertFloat64Within(t, name+".X()", expected.x, tolerance.x, actual.X())
	assertFloat64Within(t, name+".Y()", expected.y, tolerance.y, actual.Y())
}

type testCase interface {
	run(name string, t *testing.T)
}

func runTestCases(t *testing.T, testCases map[string]testCase) {
	for name, testCase := range testCases {
		t.Run(
			name,
			func(t *testing.T) {
				testCase.run(name, t)
			},
		)
	}
}
