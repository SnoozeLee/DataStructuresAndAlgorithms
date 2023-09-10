package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)

	count := 0

	for _, c := range str {
		fmt.Print(string(c))
		count++
		if count == 8 {
			count = 0
			fmt.Println()
		}
	}
	if count != 0 {
		for i := count; i < 8; i++ {
			fmt.Print("0")
		}
	}
}
