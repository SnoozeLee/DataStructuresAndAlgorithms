package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := &ListNode{1, nil}
	t := n1
	for i := 2; i < 4; i++ {
		t.Next = &ListNode{i, nil}
		t = t.Next
	}
	n2 := &ListNode{4, nil}
	t = n2
	for i := 5; i < 8; i++ {
		t.Next = &ListNode{i, nil}
		t = t.Next
	}
	s := mergeKLists([]*ListNode{n1, n2})
	for s != nil {
		fmt.Println(s.Val)
		s = s.Next
	}
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param lists ListNode类一维数组
 * @return ListNode类
 */
func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)

	var startNode *ListNode
	var tNode *ListNode

	for {
		minIdx := -1
		for i := 0; i < k; i++ {
			if lists[i] != nil {
				minIdx = i
				break
			}
		}
		for i := minIdx + 1; i < k; i++ {
			if lists[i] != nil && lists[i].Val < lists[minIdx].Val {
				minIdx = i
			}
		}
		if minIdx == -1 {
			//	lists 中没有节点了，可以返回
			return startNode
		} else {
			if tNode == nil {
				startNode = lists[minIdx]
				lists[minIdx] = lists[minIdx].Next
				tNode = startNode
				continue
			}
			tNode.Next = lists[minIdx]
			tNode = tNode.Next
			lists[minIdx] = lists[minIdx].Next
		}
	}
}
