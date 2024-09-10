package main

import (
	"testing"
)

func TestSanitiseInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "  h i mom     ",
			expected: []string{"h", "i", "mom"},
		},
	}

	for _, c := range testCases {
		actual := sanitiseInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len doesn't match: %q vs %q", actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("sanitised(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
