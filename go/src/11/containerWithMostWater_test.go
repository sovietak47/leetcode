package containerWithMostWater

import (
	"testing"
	"reflect"
)

func TestMaxArea(t *testing.T){
	type testCase struct {
		name string
		input []int
		expect int
	}

	units := []testCase{
		{
			"empty",
			[]int{},
			0,
		},
		{
			"example 1",
			[]int{1,8,6,2,5,4,8,3,7},
			49,
		},
	}

	for _, unit := range units {
		t.Run(unit.name, func(t *testing.T) {
			if actural := maxArea(unit.input); !reflect.DeepEqual(actural, unit.expect) {
				t.Errorf("maxArea failed: input: [%v], expected: [%v], actually: [%v]", unit.input, unit.expect, actural)
			}
		})
	}

}
