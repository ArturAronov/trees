package test

import (
	"testing"
	pre_order "trees-go/dfs/pre-order"
)

func TestDSFPreOrder(t *testing.T) {
	expectedResult1 := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	expectedResult2 := []int{20, 10, 5, 7, 15, 50, 30, 29, 21, 45, 49}

	outputResult1 := pre_order.PreOrderSearch(tree1)
	outputResult2 := pre_order.PreOrderSearch(tree2)

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
