package affinematrix2d_test

import (
	"github.com/GodsBoss/go-pkg/affinematrix2d"

	"testing"
)

func TestTranslation(t *testing.T) {
	runTestCases(
		t,
		map[string]testCase{
			"left_up": translationTestCase{
				translation: newVector(affinematrix2d.Left*2.0, affinematrix2d.Up*4.0),
				input:       newVector(8.5, 3.4),
				expected:    newVector(6.5, -0.6),
				tolerance:   sameTolerance(0.0001),
			},
			"left_down": translationTestCase{
				translation: newVector(affinematrix2d.Left*3.5, affinematrix2d.Down*8.0),
				input:       newVector(2.5, 1.0),
				expected:    newVector(-1.0, 9.0),
				tolerance:   sameTolerance(0.0001),
			},
			"right_up": translationTestCase{
				translation: newVector(affinematrix2d.Right*5.0, affinematrix2d.Up*3.5),
				input:       newVector(-2.5, 7.5),
				expected:    newVector(2.5, 4.0),
				tolerance:   sameTolerance(0.0001),
			},
			"right_down": translationTestCase{
				translation: newVector(affinematrix2d.Right*1.5, affinematrix2d.Down*2.5),
				input:       newVector(3.0, -1.0),
				expected:    newVector(4.5, 1.5),
				tolerance:   sameTolerance(0.0001),
			},
		},
	)
}

type translationTestCase struct {
	translation testVector
	input       testVector
	expected    testVector
	tolerance   testVector
}

func (testCase translationTestCase) run(name string, t *testing.T) {
	t.Run(
		"cartesian",
		func(t *testing.T) {
			input := affinematrix2d.VectorFromCartesian(testCase.input.x, testCase.input.y)
			translation := affinematrix2d.Translation(testCase.translation.x, testCase.translation.y)
			actual := translation.Transform(input)
			assertVectorWithin(t, name, testCase.expected, testCase.tolerance, actual)
		},
	)
	t.Run(
		"vector",
		func(t *testing.T) {
			input := affinematrix2d.VectorFromCartesian(testCase.input.x, testCase.input.y)
			translation := affinematrix2d.TranslationByVector(affinematrix2d.VectorFromCartesian(testCase.translation.x, testCase.translation.y))
			actual := translation.Transform(input)
			assertVectorWithin(t, name, testCase.expected, testCase.tolerance, actual)
		},
	)
}
