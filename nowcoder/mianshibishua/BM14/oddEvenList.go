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
 * @return ListNode类
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		// 0个元素、1个元素、2个元素的时候都不用重排
		return head
	}
	//	以下标为计算标准，0起步
	evenStart := head
	evenEnd := head
	oddStart := head.Next
	oddEnd := head.Next
	t := oddEnd.Next
	no := 2 // t 对应的下标
	for t != nil {
		if no&1 == 0 {
			//	这是出现在偶数坐标上的数，要追加到evenEnd的后面
			evenEnd.Next = t
			t = t.Next
			no++
			evenEnd = evenEnd.Next
			evenEnd.Next = oddStart // 维护连接odd链表
		} else {
			oddEnd.Next = t
			t = t.Next
			no++
			oddEnd = oddEnd.Next
			oddEnd.Next = nil // 维护末尾是nil
		}
	}
	return evenStart
}
