package houseRobberIII

//import "math"
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func rob(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left == nil && root.Right == nil {
//		return root.Val
//	}
//
//	return maxSumOfSubTree(root)
//}
//
//func maxSumOfSubTree(root *TreeNode) int {
//	return int(math.Max(float64(sumWithRoot(root)), float64(sumWithoutRoot(root))))
//}
//
//func sumWithRoot(root *TreeNode) int {
//	sum := root.Val
//	if root.Left != nil {
//		sum += sumWithoutRoot(root.Left)
//	}
//	if root.Right != nil {
//		sum += sumWithoutRoot(root.Right)
//	}
//	return sum
//}
//
//func sumWithoutRoot(root *TreeNode) int {
//	sum := 0
//	if root.Left != nil {
//		sum += maxSumOfSubTree(root.Left)
//	}
//
//	if root.Right != nil {
//		sum += maxSumOfSubTree(root.Right)
//	}
//	return sum
//}
//
