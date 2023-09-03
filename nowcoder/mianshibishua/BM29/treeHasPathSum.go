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
 * @param sum int整型
 * @return bool布尔型
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sumStart := 0
	return detect(root, sumStart, sum)
}

// detect 判断当前节点往下有没有符合题意的路径。注意：detect不判断nil节点，即需要在detect前就确保节点存在
func detect(node *TreeNode, sumBefore, flag int) bool {
	sum := sumBefore + node.Val
	if isLeaf(node) {
		if sum == flag {
			return true
		} else {
			return false
		}
	} else {
		if node.Left != nil {
			res := detect(node.Left, sum, flag)
			if res == true {
				return true
			}
		}
		if node.Right != nil {
			res := detect(node.Right, sum, flag)
			if res == true {
				return true
			}
		}
	}
	return false
}

func isLeaf(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}
