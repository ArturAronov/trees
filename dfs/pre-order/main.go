package pre_order

import "trees-go/types"

type Node = types.Node[int]

func walk(current *Node, path []int) []int {
	if current == nil {
		return path
	}

	path = append(path, current.Value)
	path = walk(current.Left, path)
	path = walk(current.Right, path)

	return path
}

func PreOrderSearch(head Node) []int {
	path := []int{}
	return walk(&head, path)
}
