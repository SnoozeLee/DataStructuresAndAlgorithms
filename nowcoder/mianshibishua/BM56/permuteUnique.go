package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{2, 0, 0, 1, 1, 3, 3}))
}

type Res struct {
	arr [][]int
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param num int整型一维数组
 * @return int整型二维数组
 */
func permuteUnique(num []int) [][]int {
	res := &Res{}
	calcPermute([]int{}, num, res)
	sort.Slice(res.arr, func(i, j int) bool {
		if len(res.arr) == 0 {
			return true
		}
		for k := 0; k < len(res.arr[0]); k++ {
			if res.arr[i][k] < res.arr[j][k] {
				return true
			} else if res.arr[i][k] > res.arr[j][k] {
				return false
			}
		}
		// 实际到不了这一步，因为任意两项一定不同
		return false
	})
	return res.arr
}

func calcPermute(head []int, tail []int, res *Res) {
	if len(tail) == 0 {
		res.arr = append(res.arr, head)
	}
	nextOneMap := make(map[int]int) // key tail上的值， value-tail上值对应的坐标
	for i := 0; i < len(tail); i++ {
		if _, exist := nextOneMap[tail[i]]; !exist {
			nextOneMap[tail[i]] = i
		}
	}
	for key, tailIdx := range nextOneMap {
		newHead := make([]int, 0, len(head)+1)
		newHead = append(newHead, head...)
		newHead = append(newHead, key)
		newTail := make([]int, 0, len(tail)-1)
		newTail = append(newTail, tail[0:tailIdx]...)
		newTail = append(newTail, tail[tailIdx+1:]...)
		calcPermute(newHead, newTail, res)
	}
}
