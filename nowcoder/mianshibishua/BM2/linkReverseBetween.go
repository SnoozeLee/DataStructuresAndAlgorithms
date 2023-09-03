package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//	找到m位置节点的前一个节点
	node := head
	idx := 1 // idx始终标记node的位置

	var mBefore *ListNode // m_before = m-1节点，将来作为局部反转链表首节点的前一节点
	if m == 1 {
		mBefore = nil
		//	node 此时恰好为m节点
	} else {
		for idx < m-1 {
			node = node.Next
			idx++
		}
		mBefore = node
		node = node.Next // node = m节点
		idx++
	}

	betweenTail := node // 将来作为局部反转链表的末尾节点

	// 局部反转
	var before *ListNode = nil
	for idx <= n {
		next := node.Next
		node.Next = before
		before = node
		node = next
		idx++
	}
	// 此时before=局部反转链表的首节点

	// 此时node指向n+1节点
	// 拼接
	betweenTail.Next = node
	if mBefore == nil {
		return before
	}
	mBefore.Next = before
	return head
}
