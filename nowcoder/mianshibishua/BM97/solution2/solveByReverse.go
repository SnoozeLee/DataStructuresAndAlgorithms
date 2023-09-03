package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := make([]int, 0)
	fmt.Println(solve(6, 2, a))
	fmt.Println(solve(6, 0, a))
	fmt.Println(solve(8, 4, b))
	fmt.Println(solve(8, 8, b))
	fmt.Println(solve(8, 12, b))
	fmt.Println(solve(0, 0, c))
}

// 这次的代码就带有技巧性了，而不是简单模拟

// 通过反转数组的方法来快速实现右移效果

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 旋转数组
 * @param n int整型 数组长度
 * @param m int整型 右移距离
 * @param a int整型一维数组 给定数组
 * @return int整型一维数组
 */
func solve(n int, m int, a []int) []int {
	if n == 0 {
		return a
	}
	m = m % n
	if m == 0 {
		return a
	}
	reverse(a, 0, n-1)
	reverse(a, 0, m-1)
	reverse(a, m, n-1)
	return a
}

// reverse 从L到R(包含L和R)的数进行反转
func reverse(a []int, L, R int) {
	if L-R >= 0 {
		// L - R 是决不能出现的，L = R 是不用计算的
		return
	}
	var t int
	for L < R {
		t = a[L]
		a[L] = a[R]
		a[R] = t
		L++
		R--
	}
}
