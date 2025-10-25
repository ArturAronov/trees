package test

import (
	"testing"
	"trees-go/dfs/in_order"
)

func TestDSFInOrder(t *testing.T) {
	expectedResult1 := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	expectedResult2 := []int{5, 7, 10, 15, 20, 21, 29, 30, 45, 49, 50}

	outputResult1 := in_order.InOrderSearch(tree1)
	outputResult2 := in_order.InOrderSearch(tree2)

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
