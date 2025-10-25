package test

import (
	"testing"
	post_order "trees-go/dfs/post-order"
)

func TestDSFPostOrder(t *testing.T) {
	expectedResult1 := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	expectedResult2 := []int{7, 5, 15, 10, 21, 29, 49, 45, 30, 50, 20}

	outputResult1 := post_order.PostOrderSearch(tree1)
	outputResult2 := post_order.PostOrderSearch(tree2)

	if len(expectedResult1) != len(outputResult1) {
		t.Errorf("Wrong output length, expected %d, recieved %d", len(expectedResult1), len(outputResult1))
	}

	if len(expectedResult2) != len(outputResult2) {
		t.Errorf("Wrong output length, expected %d, recieved %d", len(expectedResult2), len(outputResult2))
	}

	for i := range expectedResult1 {
		if outputResult1[i] != expectedResult1[i] {
			t.Errorf("Wrong output, expected %v, recieved %v", expectedResult1, outputResult1)
			break
		}
	}

	for i := range expectedResult2 {
		if outputResult2[i] != expectedResult2[i] {
			t.Errorf("Wrong output, expected %v, recieved %v", expectedResult2, outputResult2)
			break
		}
	}
}
