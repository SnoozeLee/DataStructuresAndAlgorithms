package main

import "errors"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */

func reverseBetween(head *ListNode, len int) (revHead, revTail, next *ListNode, err error) {
	//	首先检测长度够不够
	now := head
	idx := 1
	for idx <= len {
		if now == nil {
			// 不够长
			return nil, nil, nil, errors.New("len err")
		}
		now = now.Next
		idx++
	}

	//	开始反转
	var before *ListNode = nil
	idx = 1
	now = head
	for idx <= len {
		next := now.Next
		now.Next = before
		before = now
		now = next
		idx++
	}
	return before, head, now, nil
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	now := head

	var okLinkHead *ListNode = nil
	var lastTimeRevT *ListNode = nil
	firstReverse := true

	for {
		revH, revT, next, err := reverseBetween(now, k)

		if err != nil {
			if firstReverse {
				// 如果是第一次排列就长度不足，直接返回头结点
				return head
			}
			// 至少反转过一次时，返回反转链表的起始
			lastTimeRevT.Next = now
			return okLinkHead
		}

		if firstReverse {
			okLinkHead = revH
			firstReverse = false
		} else {
			lastTimeRevT.Next = revH
		}

		lastTimeRevT = revT
		now = next
	}

}
