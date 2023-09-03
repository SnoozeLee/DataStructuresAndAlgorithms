package main

import (
	"fmt"
	"math/rand"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */
func GetLeastNumbers_Solution(input []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	res := input[0:k]
	quicksort(res, 0, k-1)
	for i := k; i < len(input); i++ {
		if input[i] < res[k-1] {
			res = append(res, input[i])
			quicksort(res, 0, k)
			res = res[0:k]
		}
	}
	return input[0:k]
}

// 根据测试结果，这不知为何在牛客案例测试中，这个快排并没有之前的插入算法快
func quicksort(a []int, L, R int) {
	if R-L < 1 {
		return
	}
	randIdx := rand.Intn(R-L+1) + L
	swap(&a[randIdx], &a[R])
	i, j := L, R
	for {
		// 以a[R]为基准时，从左边找起（原因可试着分析长度为2的数组）
		for i < j && a[i] < a[R] {
			// i 的左侧都是小于a[R]的数
			i++
		}
		for i < j && a[j] >= a[R] {
			// j 的右侧都是大于等于a[R]的数
			j--
		}
		if i < j {
			swap(&a[i], &a[j])
		} else {
			// 因为i和j都是一位一位探测的，所以此时必然是 i == j
			break
		}
	}
	swap(&a[R], &a[i])
	quicksort(a, L, i-1)
	quicksort(a, i+1, R)
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	a := []int{4, 5, 1, 6, 2, 7, 3, 8, 3, 67, 3}
	fmt.Println(GetLeastNumbers_Solution(a, 4))
}
