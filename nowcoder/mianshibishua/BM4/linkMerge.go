package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	beforeResHead := &ListNode{}
	t := beforeResHead
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			t.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			t.Next = pHead2
			pHead2 = pHead2.Next
		}
		t = t.Next
	}
	if pHead1 != nil {
		t.Next = pHead1
	} else {
		t.Next = pHead2
	}
	return beforeResHead.Next
}
