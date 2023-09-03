package main

import (
	"fmt"
)

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

// 这道题目没有保证所有节点的值均不相同，也没有说 “=” 应该分布在哪一侧
// 不过按照它的文字表述，我把“=”也作错误处理就好了

func isValidBST(root *TreeNode) bool {
	a := inorderTraversal(root)
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

// 中序遍历-如果是二叉搜索树的话必然可以得到排好序的切片，如果不是，则必然得出未排序切片
func inorderTraversal(root *TreeNode) []int {
	var arr []int
	arr = inorderTraversalAssit(root, arr)
	return arr
}
func inorderTraversalAssit(node *TreeNode, arr []int) []int {
	if node.Left != nil {
		arr = inorderTraversalAssit(node.Left, arr)
	}
	arr = append(arr, node.Val)
	if node.Right != nil {
		arr = inorderTraversalAssit(node.Right, arr)
	}
	return arr
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(isValidBST(root))
}
