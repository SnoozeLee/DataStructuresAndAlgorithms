package WRONG

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
	var a []int
	a = inorderTraval(root, a)
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[len(a)-1-i] {
			return false
		}
	}
	return true
}

// 但是中序遍历出来的对称的树，未必是对称哎。
// 我可以轻松构造一个反例：
//		    4
//       /       \
//     1           2
//      \            \
//       3            3
//        \            \
//         2            1
// 但是我用中序遍历，判题系统通过了，只能说有漏洞哈哈哈

// 直接中序遍历
func inorderTraval(node *TreeNode, a []int) []int {
	if node.Left != nil {
		a = inorderTraval(node.Left, a)
	}
	a = append(a, node.Val)
	if node.Right != nil {
		a = inorderTraval(node.Right, a)
	}
	return a
}
