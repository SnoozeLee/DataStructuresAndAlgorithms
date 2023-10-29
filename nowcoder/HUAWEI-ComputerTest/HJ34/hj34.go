package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	n := len(input)
	res := make([]int32, n)

	for i, c := range input {
		res[i] = c
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] <= res[j]
	})

	for _, c := range res {
		fmt.Printf("%c", c)
	}

}
