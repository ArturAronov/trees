package test

import "trees-go/types"

type intNode = types.Node[int]

var tree1 = intNode{
	Value: 20,
	Right: &intNode{
		Value: 50,
		Right: &intNode{
			Value: 100,
			Right: nil,
			Left:  nil,
		},
		Left: &intNode{
			Value: 30,
			Right: &intNode{
				Value: 45,
				Right: nil,
				Left:  nil,
			},
			Left: &intNode{
				Value: 29,
				Right: nil,
				Left:  nil,
			},
		},
	},
	Left: &intNode{
		Value: 10,
		Right: &intNode{
			Value: 15,
			Right: nil,
			Left:  nil,
		},
		Left: &intNode{
			Value: 5,
			Right: &intNode{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}

var tree2 = intNode{
	Value: 20,
	Right: &intNode{
		Value: 50,
		Right: nil,
		Left: &intNode{
			Value: 30,
			Right: &intNode{
				Value: 45,
				Right: &intNode{
					Value: 49,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
			Left: &intNode{
				Value: 29,
				Right: nil,
				Left: &intNode{Value: 21,
					Right: nil,
					Left:  nil,
				},
			},
		},
	},
	Left: &intNode{
		Value: 10,
		Right: &intNode{
			Value: 15,
			Right: nil,
			Left:  nil,
		},
		Left: &intNode{
			Value: 5,
			Right: &intNode{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}
