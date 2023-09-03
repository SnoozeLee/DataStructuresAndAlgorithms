package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param root TreeNode类
 * @return int整型一维数组
 */
func inorderTraversal(root *TreeNode) []int {
	// write code here
	var res []int
	res = inorderTravelsalAssit(root, res)
	return res
}

func inorderTravelsalAssit(n *TreeNode, res []int) []int {
	if n == nil {
		return res
	}
	res = inorderTravelsalAssit(n.Left, res)
	res = append(res, n.Val)
	res = inorderTravelsalAssit(n.Right, res)
	return res
}
