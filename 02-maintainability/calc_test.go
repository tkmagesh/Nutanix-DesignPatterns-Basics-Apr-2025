package main

import "testing"

func Test_add(t *testing.T) {
	// arrange
	x := 100
	y := 200
	expectedResult := 300

	// act
	actualResult := add(x, y)

	// assert
	if actualResult != expectedResult {
		t.Errorf("Expected : %d, Actual : %d", expectedResult, actualResult)
	}
}
