package main

import "fmt"

func main() {
	a := []int{1, 3, 4, 6, 8, 2, 1, 5, 7, 0}
	res := findKth(a, len(a), 3)

	fmt.Println(res, "--|--", a)
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth(a []int, n int, K int) int {
	quicksort(a, 0, n-1)
	return a[n-K]
}

func quicksort(a []int, L, R int) {
	if L >= R {
		return
	}
	i := L
	j := R
	for {
		for i < j && a[i] <= a[R] {
			i++
		}
		for i < j && a[j] >= a[R] {
			j--
		}
		if i < j {
			swap(&a[i], &a[j])
		} else {
			break
		}
	}
	swap(&a[i], &a[R])
	quicksort(a, L, i-1)
	quicksort(a, i+1, R)

}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
