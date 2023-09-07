package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @return bool布尔型 有环为true
 */
func hasCycle(head *ListNode) bool {
	t1 := head
	if t1 == nil || t1.Next == nil {
		return false
	}
	t2 := head.Next
	for t2 != nil {
		// 快慢指针
		if t1 == t2 {
			return true
		}
		if t2 != nil {
			t2 = t2.Next
		}
		if t2 != nil {
			t2 = t2.Next
		}

		t1 = t1.Next // t2一直走在t1前面，且一开始就超前t1一位，所以 t2 存在时，t1的下一个必定存在
	}
	return false
}
