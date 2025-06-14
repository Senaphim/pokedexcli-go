package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "HellO WOrld",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello\n",
			expected: []string{"hello"},
		},
		{
			input:    "hello    world",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Test for matching length of output array
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match for input %s: %v != %v",
				c.input, len(actual), len(c.expected))
			continue
		}

		// Test if words in array match
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v",
					c.input, actual, c.expected)
			}
		}
	}
}
