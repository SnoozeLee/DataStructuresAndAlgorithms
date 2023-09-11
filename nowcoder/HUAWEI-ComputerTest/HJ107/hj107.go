package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	num, _ := strconv.ParseFloat(input, 64)

	fmt.Printf("%.1f\n", calcCubicRoot(num))
}

// calcCubicRoot 计算立方根，需要题目保证 |num| < 20
func calcCubicRoot(num float64) float64 {
	minusFlag := false
	if num < 0 {
		num = -num
		minusFlag = true
	}

	// 立方根一定在0-3之间
	// 计算到根的立方与num差值不超过0.001，认为是计算出来了

	res := find(num, 0, 3)
	if minusFlag {
		return -res
	}
	return res
}

// find 寻找到根的立方与num差值不超过0.001，认为是计算出来了
func find(num, L, R float64) float64 {
	m := (L + R) / 2
	var res float64
	if m*m*m > num {
		if m*m*m-num < 0.001 {
			return m
		}
		res = find(num, L, m)
	} else if m*m*m < num {
		if num-m*m*m < 0.001 {
			return m
		}
		res = find(num, m, R)
	} else {
		return m
	}
	return res
}
