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
 * @param pRoot TreeNode类
 * @return int整型二维数组
 */
func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return [][]int{}
	}

	nodes := []*TreeNode{pRoot}
	var levels [][]int
	LFlag := true
	for len(nodes) > 0 {
		level, nextNodes := DealZLevel(nodes, LFlag)
		LFlag = !LFlag
		levels = append(levels, level)
		nodes = nextNodes
	}
	return levels

}

func DealZLevel(nodes []*TreeNode, LStartFlag bool) (res []int, nextLevelNodes []*TreeNode) {
	for i := 0; i < len(nodes); i++ {
		res = append(res, nodes[i].Val)
		if nodes[i].Left != nil {
			nextLevelNodes = append(nextLevelNodes, nodes[i].Left)
		}
		if nodes[i].Right != nil {
			nextLevelNodes = append(nextLevelNodes, nodes[i].Right)
		}
	}
	if LStartFlag {
		return res, nextLevelNodes
	} else {
		return reverseArr(res), nextLevelNodes
	}
}

func reverseArr(a []int) []int {
	for i := 0; i < len(a)/2; i++ {
		t := a[i]
		a[i] = a[len(a)-i-1]
		a[len(a)-i-1] = t
	}
	return a
}
