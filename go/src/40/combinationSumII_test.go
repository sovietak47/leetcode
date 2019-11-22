package combinationSumII

import (
	"testing"
	"reflect"
)

func TestCombinationSum2(t *testing.T){
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
			"not enough candidates",
			[]int{2, 3, 4},
			10,
			nil,
		},
		{
			"example 1",
			[]int{10,1,2,7,6,1,5},
			8,
			[][]int{{7,1},{6,2},{6,1,1},{5,2,1}},
		},
		{
			"example 2",
			[]int{2,5,2,1,2},
			5,
			[][]int{{5},{2,2,1}},
		},
		{
			"example 3",
			[]int{3,5,7},
			8,
			[][]int{{5,3}},
		},
	}

	for _, unit := range units {
		t.Run(unit.name, func(t *testing.T) {
			if actural := combinationSum2(unit.candidates, unit.target); !reflect.DeepEqual(actural, unit.expect) {
				t.Errorf("combinationSum failed: input: [%v, %v], expected: [%v], actually: [%v]", unit.candidates, unit.target, unit.expect, actural)
			}
		})
	}

}
