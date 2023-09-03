package main

import "fmt"

func main() {
	permute([]int{1, 2, 3})
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param num int整型一维数组
 * @return int整型二维数组
 */
func permute(num []int) [][]int {
	// write code here
	head := make([]int, 0)
	tail := num
	res := make([][]int, 0)
	res = deal(head, tail, res)
	fmt.Println(res)
	return res
}

func deal(head, tail []int, res [][]int) [][]int {
	if len(tail) == 1 {
		res = append(res, append(head, tail[0]))
		return res
	}
	//fmt.Println("head:", head)
	//fmt.Println("tail:", tail)
	for i := 0; i < len(tail); i++ {
		var newHead []int
		var newTail []int
		newHead = append(newHead, head...)
		newHead = append(newHead, tail[i])
		newTail = append(newTail, tail[0:i]...)
		newTail = append(newTail, tail[i+1:]...)
		res = deal(newHead, newTail, res)
	}
	return res
}
