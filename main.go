package main

import (
	"trees-go/dfs/in_order"
	"trees-go/utils"
)

func main() {
	node := utils.Add([]int{6, 4, 7, 3, 2, 8, 2, 35, 9, 3, 11, 75, 32, 1, 33})
	in_order.InOrderSearch(*node)
}
