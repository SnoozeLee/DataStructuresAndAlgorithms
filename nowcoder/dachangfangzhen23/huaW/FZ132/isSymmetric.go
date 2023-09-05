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
 * @return bool布尔型
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(L, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}
	// 此时必有一个不为空
	if L == nil || R == nil {
		return false
	}
	// 此时两个都不为空
	if L.Val != R.Val {
		return false
	}
	return compare(L.Left, R.Right) && compare(L.Right, R.Left)
}
