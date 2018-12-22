package main

import "testing"

func TestGetPowerLevel(t *testing.T) {
	tests := []struct {
		input         [3]int // x, y, serial
		expectedValue int
	}{
		{[3]int{3, 5, 8}, 4},
		{[3]int{122, 79, 57}, -5},
		{[3]int{217, 196, 39}, 0},
		{[3]int{101, 153, 71}, 4},
	}
	for _, test := range tests {
		actual := getPowerLevel(test.input[0], test.input[1], test.input[2])
		if actual != test.expectedValue {
			t.Errorf("expected %d, got %d", test.expectedValue, actual)
		}
	}
}
