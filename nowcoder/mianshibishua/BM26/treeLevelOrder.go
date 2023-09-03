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
 * @return int整型二维数组
 */
func levelOrder(root *TreeNode) [][]int {
	// write code here
	// 处理没有输入的特殊情况
	if root == nil {
		return [][]int{}
	}

	nodes := []*TreeNode{root}
	// 请忽略ide报的warning： Empty slice declaration using a literal
	var levels [][]int = [][]int{}
	for len(nodes) > 0 {
		level, nextNodes := dealLevel(nodes)
		levels = append(levels, level)
		nodes = nextNodes
	}
	return levels
}

// 以后的改进，可以让nextLevelNodes提前申请好空间，减少append时多次申请空间以及拷贝的损耗
func dealLevel(levelNodes []*TreeNode) (res []int, nextLevelNodes []*TreeNode) {
	nextLevelNodes = []*TreeNode{}
	for i := 0; i < len(levelNodes); i++ {
		// 把下一层的节点装到nextLevelNodes里面
		if levelNodes[i].Left != nil {
			nextLevelNodes = append(nextLevelNodes, levelNodes[i].Left)
		}
		if levelNodes[i].Right != nil {
			nextLevelNodes = append(nextLevelNodes, levelNodes[i].Right)
		}
		//	把当前节点的数存到res里
		res = append(res, levelNodes[i].Val)
	}
	return res, nextLevelNodes
}
