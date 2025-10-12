package post_order

import "trees-go/types"

type Node = types.Node[int]

func walk(current *Node, path []int) []int {
	if current == nil {
		return path
	}

	walk(current.Left, path)
	walk(current.Right, path)
	path = append(path, current.Value)

	return path
}

func PostOrderSearch(head Node) []int {
	path := []int{}
	return walk(&head, path)
}
