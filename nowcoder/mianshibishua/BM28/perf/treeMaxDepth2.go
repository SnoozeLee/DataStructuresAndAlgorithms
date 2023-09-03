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
 * @return int整型
 */
func maxDepth(root *TreeNode) int {
	return DetectHeight(root)
}

// DetectHeight 探测某一节点树的高度，如果子树高度为零(即没有孩子),则当前节点树的高度为1
func DetectHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	//var leftDepth, rightDepth int
	// 探测子树的高度
	//if node.Left != nil {
	//	leftDepth = DetectHeight(node.Left)
	//}
	//if node.Right != nil {
	//	rightDepth = DetectHeight(node.Right)
	//}

	// 实际没有必要判断是不是空节点
	leftDepth := DetectHeight(node.Left)
	rightDepth := DetectHeight(node.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	// 当前节点树高度
	return rightDepth + 1
}
