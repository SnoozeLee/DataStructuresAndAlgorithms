package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solve("172.2.34.1"))
	fmt.Println(solve("172.2.034.1"))
	fmt.Println(solve("170.2.034.1"))
	fmt.Println(solve("170.2.a34.1"))
	fmt.Println(solve("170.2..1"))
	fmt.Println("judge v6")
	fmt.Println(solve("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(solve("2001:0db8:85a3:0:0:8a2e:0370:7334"))
	fmt.Println(solve("2001:0db8:85a3::0:8a2e:0370:7334"))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 验证IP地址
 * @param IP string字符串 一个IP地址字符串
 * @return string字符串
 */
func solve(IP string) string {
	if isIPV4(IP) {
		return "IPv4"
	}
	if isIPV6(IP) {
		return "IPv6"
	}
	return "Neither"
}

func isIPV4(IP string) bool {
	subStrList := strings.Split(IP, ".")
	// 首先可以分成4段
	if len(subStrList) != 4 {
		return false
	}
	for i := 0; i < 4; i++ {
		subStr := make([]int32, 0)
		for _, char := range subStrList[i] {
			subStr = append(subStr, char)
		}
		valid, _ := calcIPV4Num(subStr)
		if valid == false {
			return false
		}
	}
	return true
}

func calcIPV4Num(s []int32) (isValid bool, res int) {
	// 首先保证只能一到三位数
	if len(s) > 3 || len(s) == 0 {
		return false, -1
	}
	// 第一个不能为0
	if s[0] == '0' {
		return false, -1
	}
	//	接着保证每个数大于等于'0'小于等于'9'
	res = 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			res = res*10 + int(s[i]-'0')
		} else {
			return false, -1
		}
	}
	// 控制上限
	if res > 255 {
		return false, -1
	}
	return true, res
}

func isIPV6(IP string) bool {
	subStrList := strings.Split(IP, ":")
	// 先判断组数合不合要求
	if len(subStrList) != 8 {
		return false
	}
	// 再分组判断各组字符串合不合要求
	for i := 0; i < 8; i++ {
		subStr := make([]int32, 0)
		for _, char := range subStrList[i] {
			subStr = append(subStr, char)
		}

		if valid := calcIPV6NumValid(subStr); valid == false {
			return false
		}
	}
	//	各组判断结束
	return true
}

func calcIPV6NumValid(s []int32) (isValid bool) {
	// 只能是1到4个数
	if len(s) > 4 || len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		t := s[i]
		if (t >= '0' && t <= '9') || (t >= 'a' && t <= 'f') || (t >= 'A' && t <= 'F') {

		} else {
			return false
		}
	}
	return true
}
