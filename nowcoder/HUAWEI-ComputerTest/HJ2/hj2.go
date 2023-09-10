package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	scanner.Scan()
	aimStr := scanner.Text()
	var aim int32
	for _, c := range aimStr {
		aim = c
	}
	count := 0
	var aim2 int32
	hasAim2 := false
	if aim >= 'a' && aim <= 'z' {
		aim2 = aim + 'A' - 'a'
		hasAim2 = true
	} else if aim >= 'A' && aim <= 'Z' {
		aim2 = aim + 'a' - 'A'
		hasAim2 = true
	}
	for _, c := range line {
		if c == aim || (hasAim2 && c == aim2) {
			count++
		}
	}
	fmt.Println(count)
}
