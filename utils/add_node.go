package utils

import (
	"fmt"
	"trees-go/types"
)

type Node = types.Node[int]

func AddNode(head *Node, newNode *Node) *Node {
	if head == nil {
		return newNode
	}

	if head.Value > newNode.Value {
		head.Left = AddNode(head.Left, newNode)
	}

	if head.Value < newNode.Value {
		head.Right = AddNode(head.Right, newNode)
	}

	return head
}

func Add(values []int) *Node {
	var head *Node

	for i := range values {
		head = AddNode(head, &Node{
			Value: values[i],
		})
	}

	fmt.Println(*head)

	return head
}
