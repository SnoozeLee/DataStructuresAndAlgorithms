package main

import (
	"errors"
)

type stack struct {
	arr    []int
	topIdx int
}

func (s *stack) push(node int) {
	if len(s.arr) == s.topIdx {
		// 要么长度恰好是top下标
		s.arr = append(s.arr, node)
		s.topIdx = s.topIdx + 1
	} else {
		//	要么top下标小于长度，因为被pop掉了
		s.arr[s.topIdx] = node
		s.topIdx++
	}
}
func (s *stack) pop() (int, error) {
	if s.topIdx == 0 {
		return -1, errors.New("pop nothing")
	}
	s.topIdx--
	return s.arr[s.topIdx], nil
}

var s1 stack = stack{[]int{}, 0}
var s2 stack = stack{[]int{}, 0}

func Push(node int) {
	s1.push(node)
}

func Pop() int {
	res, err := s2.pop()
	if err != nil {
		//	s2为空，需要把s1的数装载过来
		for {
			t, e := s1.pop()
			if e != nil {
				//	s1中的数都装到s2中了
				break
			}
			s2.push(t)
		}
		res, err = s2.pop()
		if err != nil {
			//彻底没数了
		}
		return res
	}
	return res
}
