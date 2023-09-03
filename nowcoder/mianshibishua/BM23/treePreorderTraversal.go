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
func preorderTraversal(root *TreeNode) []int {
	// write code here
	var res []int
	res = preorderTraversalAssit(root, res)
	return res
}

func preorderTraversalAssit(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = append(res, node.Val)
	res = preorderTraversalAssit(node.Left, res)
	res = preorderTraversalAssit(node.Right, res)
	return res
}
