package bfs

import "trees-go/types"

type Node = types.Node[int]

func shift(list []Node) (*Node, []Node) {
	if len(list) < 1 {
		return nil, list
	}

	return &list[0], list[1:] //start from index 1 : until the end of array
}

func BSF(head Node, needle int) bool {
	queue := []Node{head}

	for len(queue) > 0 {
		current, list := shift(queue)

		if current.Value == needle {
			return true
		}

		if current.Left != nil {
			queue = append(list, *current.Left)
		}

		if current.Right != nil {
			queue = append(list, *current.Right)
		}
	}

	return false
}
