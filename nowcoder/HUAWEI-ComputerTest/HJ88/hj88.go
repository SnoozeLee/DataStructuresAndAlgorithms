package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pokerType int

const singleType pokerType = 1
const doubleType pokerType = 2
const trebleType pokerType = 3
const shunziType pokerType = 5
const boomType pokerType = 8

func Type(poker string) pokerType {
	p := strings.Split(poker, " ")
	switch len(p) {
	case 1:
		return singleType
	case 2:
		if p[0][0:1] == "j" || p[0][0:1] == "J" {
			return boomType
		}
		return doubleType
	case 3:
		return trebleType
	case 4:
		return boomType
	case 5:
		return shunziType
	}
	// 不存在的情况
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()
	pokers := strings.Split(input, "-")

	if Type(pokers[0]) == Type(pokers[1]) {
		if pokerValue(pokers[0]) > pokerValue(pokers[1]) {
			fmt.Println(pokers[0])
		} else {
			fmt.Println(pokers[1])
		}
		return
	}

	// 此时是不同类型的比较

	// 一个炸弹+一种普通牌比较 的可能
	if Type(pokers[0]) == boomType {
		fmt.Println(pokers[0])
		return
	} else if Type(pokers[1]) == boomType {
		fmt.Println(pokers[1])
		return
	}

	// 不可比较
	fmt.Println("ERROR")
	return
}

// 适合用来比较都是所有同类型的，包括王炸和普通炸弹的比较
func pokerValue(poker string) int {
	c := poker[0:1]
	if c >= "3" && c <= "9" {
		n, _ := strconv.ParseInt(c, 10, 64)
		return int(n)
	}
	switch c {
	case "1":
		return 10
	case "J":
		c2 := poker[1:2]
		if c2 == "O" {
			return 17
		}
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	case "2":
		return 15
	case "j":
		return 16
	}
	return -1 // 不可能出现
}
