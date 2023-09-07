package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类 the head
 * @return bool布尔型
 */

// 我觉得更合适的命名应该是 isPali 因为回文的单词是 palindrome
// 如果允许反链表(会修改链表结构)的话可以采用如下方法。如果不允许修改，也可以判断完之后再把链表反回来
func isPail(head *ListNode) bool {
	length := calcLinkLen(head)
	if length <= 1 {
		return true
	}

	var before *ListNode = nil
	mid := head
	after := head.Next

	for i := 0; i < length/2; i++ {
		mid.Next = before
		before = mid
		mid = after
		after = after.Next
	}

	if length&1 != 0 {
		//	奇数长度
		mid = mid.Next
	}

	for before != nil {
		if before.Val != mid.Val {
			return false
		}
		before = before.Next
		mid = mid.Next
	}
	return true
}

func calcLinkLen(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
