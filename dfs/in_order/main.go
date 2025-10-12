package in_order

import "trees-go/types"

type Node = types.Node[int]

func walk(current *Node, path []int) []int {
	if current == nil {
		return path
	}

	walk(current.Left, path)
	path = append(path, current.Value)
	walk(current.Right, path)

	return path
}

func InOrderSearch(head Node) []int {
	path := []int{}
	return walk(&head, path)
}
