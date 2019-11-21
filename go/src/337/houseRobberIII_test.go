package houseRobberIII

import (
	"reflect"
	"testing"
)

func sliceToTreeNode(input []int) *TreeNode {
	var root *TreeNode

	if len(input) == 0 {
		root = nil
	} else {
		root = &TreeNode{
			Val: input[0],
		}

		currentLevelNode := []*TreeNode{root}
		for i := 1; i < len(input); {
			nextLevelNode := []*TreeNode{}
			for j := range currentLevelNode {
				currentLevelNode[j].Left = &TreeNode{}
				currentLevelNode[j].Right = &TreeNode{}
				nextLevelNode = append(nextLevelNode, currentLevelNode[j].Left, currentLevelNode[j].Right)
			}

			for k := range nextLevelNode {
				if i < len(input) {
					nextLevelNode[k].Val = input[i]
				} else {
					nextLevelNode[k] = nil
				}
				i++
			}
			currentLevelNode = nextLevelNode
		}
	}
	return root
}

func TestRob(t *testing.T) {
	type testCase struct {
		name   string
		input  []int
		expect int
	}

	units := []testCase{
		{
			"empty",
			[]int{},
			0,
		},
		{
			"one node",
			[]int{3},
			3,
		},
		{
			"example 1",
			[]int{3, 2, 3, 0, 3, 0, 1},
			7,
		},
		{
			"example 2",
			[]int{3, 4, 5, 1, 3, 0, 1},
			9,
		},
	}

	for _, unit := range units {
		t.Run(unit.name, func(t *testing.T) {
			root := sliceToTreeNode(unit.input)
			if actural := rob(root); !reflect.DeepEqual(actural, unit.expect) {
				t.Errorf("rob failed: input: [%v], expected: [%v], actually: [%v]", unit.input, unit.expect, actural)
			}
		})
	}

}
