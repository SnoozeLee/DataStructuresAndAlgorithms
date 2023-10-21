package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 比较版本号
 * @param version1 string字符串
 * @param version2 string字符串
 * @return int整型
 */
func compare(version1 string, version2 string) int {
	// write code here
	v1StrArr := strings.Split(version1, ".")
	v2StrArr := strings.Split(version2, ".")

	maxLen := 1
	if len(v1StrArr) >= len(v2StrArr) {
		maxLen = len(v1StrArr)
	} else {
		maxLen = len(v2StrArr)
	}
	v1arr := make([]int64, maxLen)
	v2arr := make([]int64, maxLen)

	for i := 0; i < len(v1StrArr); i++ {
		v1arr[i], _ = strconv.ParseInt(v1StrArr[i], 10, 64)
	}

	for i := 0; i < len(v2StrArr); i++ {
		v2arr[i], _ = strconv.ParseInt(v2StrArr[i], 10, 64)
	}

	for i := 0; i < maxLen; i++ {
		fmt.Println(v1arr[i], v2arr[i])
		if v1arr[i] > v2arr[i] {
			return 1
		} else if v1arr[i] < v2arr[i] {
			return -1
		}
	}

	return 0
}
