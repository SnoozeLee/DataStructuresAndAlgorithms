package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */

// 结合插入排序的stack
type minStack struct {
	arr []int
	k   int
}

func newMinStack(arr []int, k int) *minStack {
	s := &minStack{k: k}
	s.arr = arr[:k]
	insertionSort(s.arr, k)
	return s
}

func (s *minStack) Insert(n int) {
	if s.k-1 > -1 && n < s.arr[s.k-1] {
		s.arr[s.k-1] = n
		insertionSort(s.arr, s.k)
	}
}

// 采用insertionSort不能达到要求的O(nlogk)时间复杂度
// 而且没必要每次来一个数都排一次序
func insertionSort(arr []int, len int) {
	if len <= 1 {
		return
	}
	for i := 1; i < len; i++ {
		t := arr[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			// 判断j位置上的数
			if t < arr[j] {
				// 后移一格
				arr[j+1] = arr[j]
			} else {
				// j位置的数找到不比t大
				break
			}
		}
		// j位置的数找到不比t大,在j+1位置放t
		arr[j+1] = t
	}
}

func GetLeastNumbers_Solution(input []int, k int) []int {
	s := newMinStack(input, k)
	for i := k; i < len(input); i++ {
		s.Insert(input[i])
	}
	return s.arr
}

//func main() {
//	a := []int{3, 2, 4, 5, 1}
//}
