package houseRobberIII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
  sumWithRoot, sumWithoutRoot := twoSituationsSumOfSubTree(root)    
	return max(sumWithRoot, sumWithoutRoot)
}

func twoSituationsSumOfSubTree(root *TreeNode) (sumWithSelf, sumWithoutSelf int) {
    if root == nil {
		return 0, 0
	}

    LeftSubTreeSumWithLeft, LeftSubTreeSumWithoutLeft:= twoSituationsSumOfSubTree(root.Left)
    RightSubTreeSumWithRight, RightSubTreeSumWithoutRight:= twoSituationsSumOfSubTree(root.Right)

    sumWithSelf =  root.Val + LeftSubTreeSumWithoutLeft + RightSubTreeSumWithoutRight
    sumWithoutSelf = max(LeftSubTreeSumWithLeft, LeftSubTreeSumWithoutLeft) + max(RightSubTreeSumWithRight, RightSubTreeSumWithoutRight)
    return 
}

func max(a,b int) int{
    if a> b {
        return a
    }
    return b
}
