package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	// create a slice of test case structs
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   Hello WORLD   ",
			expected: []string{"hello", "world"},
		},
	}

	// loop over the cases and run tests
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected Length: %v, Actual Length: %v", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected Word at the index: %v, Actual Word: %v", word, expectedWord)
			}
		}
	}
}
