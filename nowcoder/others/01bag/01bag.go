package main

import "fmt"

type Good struct {
	weight int
	value  int
}

func main() {
	bag := 0
	n := 0
	_, _ = fmt.Scanf("%d %d\n", &bag, &n)

	state := make([][]int, n+1)
	for i := 0; i < len(state); i++ {
		state[i] = make([]int, n+1)
	}

	goods := []Good{
		{0, 0},
	}

	for i := 1; i <= n; i++ {
		goods = append(goods, Good{0, 0})
		_, _ = fmt.Scanf("%d %d\n", &goods[i].weight, &goods[i].value)
	}

	for row := 1; row < n+1; row++ {
		for w := 1; w <= bag; w++ {
			if w < goods[row].weight {
				state[row][w] = state[row-1][w]
			} else {
				if state[row-1][w-goods[row].weight]+goods[row].value > state[row-1][w] {
					state[row][w] = state[row-1][w-goods[row].weight] + goods[row].value
				} else {
					state[row][w] = state[row-1][w]
				}
			}
		}
	}

	fmt.Println(state[len(state)-1][bag])
}
