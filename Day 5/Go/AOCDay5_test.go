package main

import "testing"

func TestReact(t *testing.T) {
	tests := []struct {
		input         []string
		expectedValue bool
	}{
		{[]string{"a", "a"}, false},
		{[]string{"a", "b"}, false},
		{[]string{"A", "A"}, false},
		{[]string{"a", "B"}, false},
		{[]string{"a", "A"}, true},
		{[]string{"A", "a"}, true},
	}
	for _, test := range tests {
		actual := react(test.input[0], test.input[1])
		if actual != test.expectedValue {
			t.Errorf("expected %t, got %t", test.expectedValue, actual)
		}
	}
}
