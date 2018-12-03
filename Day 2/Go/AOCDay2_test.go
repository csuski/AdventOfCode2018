package main

import "testing"

func TestLettersInCommon(t *testing.T) {
	tests := []struct {
		input         []string
		expectedValue string
	}{
		{[]string{"aaa", "aaa"}, "aaa"},
		{[]string{"aaa", "abb"}, "a"},
	}
	for _, test := range tests {
		actual := lettersInCommon(test.input[0], test.input[1])
		if actual != test.expectedValue {
			t.Errorf("expected %s, got %s", test.expectedValue, actual)
		}
	}
}
