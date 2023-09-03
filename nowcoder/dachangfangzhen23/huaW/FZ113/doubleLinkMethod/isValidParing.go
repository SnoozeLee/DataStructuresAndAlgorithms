package main

import "fmt"

func main() {
	str := "ACDB"
	str2 := "ABCD"
	str3 := "ACBD"
	fmt.Println(isValidPairing(str))
	fmt.Println(isValidPairing(str2))
	fmt.Println(isValidPairing(str3))
}

type Node struct {
	v      int32
	before *Node
	next   *Node
}

func isPair(nA, nB *Node) bool {
	if nA == nil || nB == nil {
		return false
	}

	switch nA.v {
	case 'A':
		if nB.v == 'B' {
			return true
		}
		return false
	//	AB有顺序要求，AB配对，但BA不配对
	//case 'B':
	//	if nB.v == 'A' {
	//		return true
	//	}
	//	return false
	case 'C':
		if nB.v == 'D' {
			return true
		}
		return false
	//case 'D':
	//	if nB.v == 'C' {
	//		return true
	//	}
	//	return false
	default:
		return false
	}
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return bool布尔型
 */
func isValidPairing(s string) bool {
	var head *Node
	var t *Node = nil
	for _, char := range s {
		newNode := &Node{v: char}
		newNode.before = t
		if t != nil {
			t.next = newNode
		} else {
			head = newNode
		}
		t = newNode
	}
	// 双向链表生成结束

	t = head
	for t != nil {
		if isPair(t, t.next) {
			var isBefore bool
			t, isBefore = deleteTwoNodes(t)
			if !isBefore {
				// 前面没有节点了，head要变化
				head = t
			}
		} else {
			t = t.next
		}
	}
	// 如果head链为空，说明都消除完毕了
	if head == nil {
		return true
	} else {
		return false
	}
}

// deleteTwoNodes 从双向链表中删除t和t.next
// 需要保证t和t.next都存在才可以删除，否则会发生意外操作
// 删除后，如果两个节点前有节点，则返回前面的节点
// 否则，返回后面的节点，如果后面的节点是空结点，那返回空结点
func deleteTwoNodes(t *Node) (nextJudge *Node, isBefore bool) {
	var tBefore *Node
	var tTail *Node
	if t.before != nil {
		tBefore = t.before
	}
	if t.next.next != nil {
		tTail = t.next.next
	}

	if tBefore != nil {
		tBefore.next = tTail
		if tTail != nil {
			tTail.before = tBefore
		}
		return tBefore, true
	} else {
		if tTail != nil {
			tTail.before = nil
		}
		return tTail, false
	}
}
