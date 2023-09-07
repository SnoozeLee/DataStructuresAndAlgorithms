package main

// 这道题题目核心代码没给ListNode的样式，我只好去BM6复制一个

//type ListNode struct {
//	Next *ListNode
//	Val  int
//}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}
	p1 := pHead
	p2 := pHead.Next

	for p2 != nil {
		if p1 == p2 {
			// 有环
			break
		}
		if p2.Next != nil && p2.Next.Next != nil {
			// 快慢指针
			p1 = p1.Next
			p2 = p2.Next.Next
		} else {
			// 无环
			return nil
		}
	}

	//	此时p1和p2在同一位置，让p2走一圈进行计数
	p2 = p2.Next
	count := 1
	for p2 != p1 {
		p2 = p2.Next
		count++
	}

	p1 = pHead
	p2 = pHead
	for i := 0; i < count; i++ {
		p2 = p2.Next
	}

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
