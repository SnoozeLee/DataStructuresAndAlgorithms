package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(calc(n))
}

func calc(n int) int {
	count := 0
	for i := 7; i <= n; i++ {
		if judge7(i) {
			count++
		}
	}
	return count
}

func judge7(n int) bool {
	// 实际题目保证了n是正整数也可以不判断0情况
	// if n == 0 {
	//     return false
	// }

	if n%7 == 0 {
		return true
	}

	for n != 0 {
		if n%10 == 7 {
			return true
		}
		n = n / 10
	}
	return false
}
