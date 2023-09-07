package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */

// 我一想到的方法还是哈希存表 。。。

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	len1 := calcLinkLen(pHead1)
	len2 := calcLinkLen(pHead2)

	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			pHead1 = pHead1.Next
		}
	} else if len2 > len1 {
		for i := 0; i < len2-len1; i++ {
			pHead2 = pHead2.Next
		}
	}

	for pHead1 != pHead2 {
		pHead1 = pHead1.Next
		pHead2 = pHead2.Next
	}
	return pHead1
}

// calcLinkLen 计算链表长度，请勿传入有环链表！程序无法处理
func calcLinkLen(pHead *ListNode) int {
	count := 0
	for pHead != nil {
		count++
		pHead = pHead.Next
	}
	return count
}
