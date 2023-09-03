package main

import (
	"math"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
 */
func LCS(str1 string, str2 string) string {
	table := make([][]int, 0)
	for i := 0; i < len(str1); i++ {
		table = append(table, make([]int, len(str2)))
	}
	// 记录最大值
	maxI := 0
	maxJ := 0
	maxV := math.MinInt
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				if i == 0 || j == 0 {
					// 左边界和上边界
					table[i][j] = 1
				} else {
					// 非边界
					table[i][j] = table[i-1][j-1] + 1
				}
				//	更新最大值
				if table[i][j] > maxV {
					maxV = table[i][j]
					maxI = i
					maxJ = j
				}

			} else {
				table[i][j] = 0
			}
		}
	}

	//	找到最长子串
	idx := maxI
	jdx := maxJ
	resStr := make([]uint8, 0)
	for idx >= 0 && jdx >= 0 && table[idx][jdx] > 0 {
		resStr = append(resStr, str1[idx])
		idx--
		jdx--

	}
	for i := 0; i < len(resStr)/2; i++ {
		t := resStr[i]
		resStr[i] = resStr[len(resStr)-1-i]
		resStr[len(resStr)-1-i] = t
	}
	return string(resStr)
}
