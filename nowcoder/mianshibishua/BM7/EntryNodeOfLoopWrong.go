package main

// 这道题题目核心代码没给ListNode的样式，我只好去BM6复制一个

type ListNode struct {
	Next *ListNode
	Val  int
}

// 我先尝试用一波空间复杂度在O(n)的做法
// 不过判题系统给我过了

func EntryNodeOfLoop_1(pHead *ListNode) *ListNode {
	listMap := make(map[*ListNode]int)
	t := pHead
	for t != nil {
		if _, ok := listMap[t]; ok {
			return t
		} else {
			listMap[t] = 1
			t = t.Next
		}
	}
	return nil
}
