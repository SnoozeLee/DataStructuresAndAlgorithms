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
func postorderTraversal(root *TreeNode) []int {
	return POTAssit(root, []int{})
}

func POTAssit(n *TreeNode, res []int) []int {
	if n == nil {
		return res
	}
	res = POTAssit(n.Left, res)
	res = POTAssit(n.Right, res)
	res = append(res, n.Val)
	return res
}
