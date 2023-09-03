package main

// stack 说明：topIdx 指向栈顶元素的下一个元素
type stack struct {
	arr    []int
	topIdx int
}

func (s *stack) push(n int) {
	if len(s.arr) == s.topIdx {
		s.arr = append(s.arr, n)
	} else {
		s.arr[s.topIdx] = n
	}
	s.topIdx++
}
func (s *stack) pop() int {
	s.topIdx--
	return s.arr[s.topIdx]
}
func (s *stack) top() int {
	return s.arr[s.topIdx-1]
}

var s stack = stack{}
var minIdxStack stack = stack{}

func Push(node int) {
	//	维护最小值下标
	if s.topIdx == 0 {
		//	刚入栈底，必定是最小数
		minIdxStack.push(0)
	} else if node < s.arr[minIdxStack.top()] {
		minIdxStack.push(s.topIdx)
	}

	s.push(node)
}
func Pop() {
	s.pop()
	// 如果最小值被弹出了，minIdxStack 需要同样弹出
	if s.topIdx == minIdxStack.top() {
		minIdxStack.pop()
	}
}
func Top() int {
	return s.top()
}

func Min() int {
	return s.arr[minIdxStack.top()]
}
