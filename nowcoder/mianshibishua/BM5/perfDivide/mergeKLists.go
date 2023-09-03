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
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	resList := lists[0]
	for i := 1; i < len(lists); i++ {
		resList = merge2Lists(resList, lists[i])
	}
	return resList
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	n1 := l1
	n2 := l2
	var startNode *ListNode
	var tNode *ListNode

	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	// 选定头结点
	if n1.Val < n2.Val {
		startNode = n1
		n1 = n1.Next
	} else {
		startNode = n2
		n2 = n2.Next
	}
	tNode = startNode

	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			tNode.Next = n1
			n1 = n1.Next
		} else {
			tNode.Next = n2
			n2 = n2.Next
		}
		tNode = tNode.Next
	}

	if n1 != nil {
		tNode.Next = n1
	} else {
		tNode.Next = n2
	}
	return startNode
}
