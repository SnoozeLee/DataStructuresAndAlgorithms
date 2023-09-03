package main

import "fmt"

func main() {
	fmt.Println(LCS("ABCDEFG", "BCDEGF"))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * longest common subsequence
 * @param s1 string字符串 the string
 * @param s2 string字符串 the string
 * @return string字符串
 */
func LCS(s1 string, s2 string) string {
	str := calcLCS(s1, s2)
	if len(str) > 0 {
		return str
	}
	return "-1"
}

type Dir int

const (
	noDir  Dir = 0
	left   Dir = 1
	up     Dir = 2
	leftup Dir = 3
)

type Cell struct {
	maxLen int // 最长公共子序列长度
	dir    Dir // 前一个cell所在的方位
}

func calcLCS(s1, s2 string) string {
	tHeight := len(s1) + 1
	tWidth := len(s2) + 1
	table := make([][]Cell, 0) // 查找表（子问题表/动态规划表）
	for i := 0; i < tHeight; i++ {
		table = append(table, make([]Cell, tWidth))
	}
	//	go自带初始化，左边界和上边界都自动初始化为0
	for j := 1; j < tWidth; j++ {
		for i := 1; i < tHeight; i++ {
			if s1[i-1] == s2[j-1] {
				table[i][j].dir = leftup
				table[i][j].maxLen = table[i-1][j-1].maxLen + 1
			} else if table[i-1][j].maxLen >= table[i][j-1].maxLen {
				table[i][j].dir = up
				table[i][j].maxLen = table[i-1][j].maxLen
			} else {
				table[i][j].dir = left
				table[i][j].maxLen = table[i][j-1].maxLen
			}
		}
	}
	//	表格已经生成，开始往回拼接答案
	resStr := make([]uint8, 0)
	i := tHeight - 1
	j := tWidth - 1
	for i > 0 && j > 0 {
		switch table[i][j].dir {
		case leftup:
			resStr = append(resStr, s1[i-1])
			i--
			j--
		case up:
			i--
		case left:
			j--
		}
	}
	for strI := 0; strI < len(resStr)/2; strI++ {
		t := resStr[strI]
		resStr[strI] = resStr[len(resStr)-1-strI]
		resStr[len(resStr)-1-strI] = t
	}
	return string(resStr)
}
