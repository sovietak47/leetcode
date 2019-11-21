package asteroidCollision

import (
	"testing"
	"reflect"
)

func TestAsteroidCollision(t *testing.T){
	type testCase struct {
		name string
		input []int
		expect []int
	}

	units := []testCase{
		{
			"all left",
			[]int{-1, -3, -5},
			[]int{-1, -3, -5},
		},
		{
			"all right",
			[]int{1, 3, 5},
			[]int{1, 3, 5},
		},
		{
			"run own way",
			[]int{-1, -2, -3, 5, 7},
			[]int{-1, -2, -3, 5, 7},
		},
		{
			"both been destroyed",
			[]int{-1, -2, -3, 5, -5},
			[]int{-1, -2, -3},
		},
		{
			"new left moved asteroid been destroyed",
			[]int{-1, -2, -3, 5, -1},
			[]int{-1, -2, -3, 5},
		},
		{
			"new left moved asteroid destroyed all right moved asteroids",
			[]int{-1, -2, -3, 5, -10},
			[]int{-1, -2, -3, -10},
		},
	}

	for _, unit := range units {
		t.Run(unit.name, func(t *testing.T) {
			if actural := asteroidCollision(unit.input); !reflect.DeepEqual(actural, unit.expect) {
				t.Errorf("asteroidCollision failed: input: [%v], expected: [%v], actually: [%v]", unit.input, unit.expect, actural)
			}
		})
	}

}
