package BM28

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
	if root == nil {
		return 0
	}
	return DetectDepth(root, 1)
}

// 下次改进：从下往上数，而不是从上往下数
func DetectDepth(node *TreeNode, depth int) int {
	// depth 当前节点的深度
	max := depth
	if node.Left != nil {
		leftDepth := DetectDepth(node.Left, depth+1)
		if leftDepth > max {
			max = leftDepth
		}
	}
	if node.Right != nil {
		rightDepth := DetectDepth(node.Right, depth+1)
		if rightDepth > max {
			max = rightDepth
		}
	}
	return max
}
