package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	hashMap := make(map[int64]bool, 0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n_str := scanner.Text()
	n_64, _ := strconv.ParseInt(n_str, 10, 64)
	n := int(n_64)

	count := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		num_str := scanner.Text()
		num, _ := strconv.ParseInt(num_str, 10, 64)
		if _, ok := hashMap[num]; ok == true {

		} else {
			hashMap[num] = true
			count++
		}
	}

	res := make([]int64, 0, count)
	for key, _ := range hashMap {
		res = append(res, key)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] <= res[j]
	})

	for i := 0; i < count; i++ {
		fmt.Println(res[i])
	}
}
