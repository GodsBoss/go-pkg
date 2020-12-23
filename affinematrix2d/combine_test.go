package affinematrix2d_test

import (
	"fmt"

	"github.com/GodsBoss/go-pkg/affinematrix2d"

	"testing"
)

func TestCombine(t *testing.T) {
	runTestCases(
		t,
		map[string]testCase{
			"identity": combineTestCase{
				transformations: make([]affinematrix2d.Transformation, 0),
				input:           newVector(2.5, 3.5),
				expected:        newVector(2.5, 3.5),
				tolerance:       sameTolerance(0.0001),
			},
			"one_fixed": combineTestCase{
				transformations: []affinematrix2d.Transformation{fixedTransformation(newVector(1.8, -2.4))},
				input:           newVector(-3.4, 5.8),
				expected:        newVector(1.8, -2.4),
				tolerance:       sameTolerance(0.0001),
			},
			"three_combined": combineTestCase{
				transformations: []affinematrix2d.Transformation{
					affinematrix2d.Translation(2.5, -3.0),
					affinematrix2d.Scale(2.0, 1.5),
					affinematrix2d.Rotation(affinematrix2d.FullAngle * 0.25 * affinematrix2d.Clockwise),
				},
				input:     newVector(3.0, 2.5),
				expected:  newVector(-2.5, 1.5),
				tolerance: sameTolerance(0.0001),
			},
			"custom_transformation_fixed_in_between": combineTestCase{
				transformations: []affinematrix2d.Transformation{
					affinematrix2d.Translation(-3.0, 4.5),
					affinematrix2d.Scale(2.0, 3.0),
					fixedTransformation(newVector(1.8, -2.4)),
					affinematrix2d.Scale(1.2, 2.3),
					affinematrix2d.Translation(-3.0, 4.1),
					affinematrix2d.Rotation(-1.2),
				},
				input:     newVector(-3.4, 5.8), // Irrelevant, because fixedTransformation ignores input
				expected:  newVector(0.6, -2.7),
				tolerance: sameTolerance(0.0001),
			},
			"custom_transformation_id_first": combineTestCase{
				transformations: []affinematrix2d.Transformation{
					affinematrix2d.Translation(-0.4, 0.8),
					testIdentity{},
				},
				input:     newVector(-1.4, 5.2),
				expected:  newVector(-1.8, 6.0),
				tolerance: sameTolerance(0.0001),
			},
			"custom_transformation_id_last": combineTestCase{
				transformations: []affinematrix2d.Transformation{
					testIdentity{},
					affinematrix2d.Translation(-0.4, 0.8),
				},
				input:     newVector(-1.4, 5.2),
				expected:  newVector(-1.8, 6.0),
				tolerance: sameTolerance(0.0001),
			},
		},
	)
}

type combineTestCase struct {
	transformations []affinematrix2d.Transformation
	input           testVector
	expected        testVector
	tolerance       testVector
}

func (testCase combineTestCase) run(name string, t *testing.T) {
	combinedTransformation := affinematrix2d.Combine(testCase.transformations...)
	fmt.Printf("combined_transformation = %#v\n", combinedTransformation)
	actual := combinedTransformation.Transform(affinematrix2d.VectorFromCartesian(testCase.input.x, testCase.input.y))
	assertVectorWithin(t, name, testCase.expected, testCase.tolerance, actual)
}

type fixedTransformation testVector

func (ft fixedTransformation) Transform(_ affinematrix2d.Vector) affinematrix2d.Vector {
	return affinematrix2d.VectorFromCartesian(ft.x, ft.y)
}

type testIdentity struct{}

func (ti testIdentity) Transform(v affinematrix2d.Vector) affinematrix2d.Vector {
	return v
}
