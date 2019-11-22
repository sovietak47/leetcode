package combinationSum

import (
	"testing"
	"reflect"
)

func TestCombinationSum(t *testing.T){
	type testCase struct {
		name string
		candidates []int
		target int
		expect [][]int
	}

	units := []testCase{
		{
			"empty",
			[]int{},
			7,
			nil,
		},
		{
			"no result",
			[]int{2, 3, 4},
			1,
			nil,
		},
		{
			"example 1",
			[]int{2,3,6,7},
			7,
			[][]int{{7},{3,2,2}},
		},
		{
			"example 2",
			[]int{2, 3, 5},
			8,
			[][]int{{5,3},{3,3,2},{2,2,2,2}},
		},
	}

	for _, unit := range units {
		t.Run(unit.name, func(t *testing.T) {
			if actural := combinationSum(unit.candidates, unit.target); !reflect.DeepEqual(actural, unit.expect) {
				t.Errorf("combinationSum failed: input: [%v, %v], expected: [%v], actually: [%v]", unit.candidates, unit.target, unit.expect, actural)
			}
		})
	}

}
