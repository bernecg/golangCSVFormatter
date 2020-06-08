package tests

import (
	"testing"

	matrix "../matrix"
)

func TestEcho(t *testing.T) {
	testMatrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := "1,2,3\n4,5,6\n7,8,9\n"
	response := matrix.Echo(testMatrix)

	if expected != response {
		t.Errorf("Echo failed, expected %v but got %v", expected, response)
	}
}

func TestInvert(t *testing.T) {
	testMatrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := "1,4,7\n2,5,8\n3,6,9\n"
	response := matrix.Invert(testMatrix)

	if expected != response {
		t.Errorf("Echo failed, expected %v but got %v", expected, response)
	}
}

func TestFlatten(t *testing.T) {
	testMatrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := "1,2,3,4,5,6,7,8,9\n"
	response := matrix.Flatten(testMatrix)

	if expected != response {
		t.Errorf("Echo failed, expected %v but got %v", expected, response)
	}
}

func TestSum(t *testing.T) {
	testMatrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := "45\n"
	response := matrix.Sum(testMatrix)

	if expected != response {
		t.Errorf("Echo failed, expected %v but got %v", expected, response)
	}
}

func TestMultiply(t *testing.T) {
	testMatrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := "362880\n"
	response := matrix.Multiply(testMatrix)

	if expected != response {
		t.Errorf("Echo failed, expected %v but got %v", expected, response)
	}
}
