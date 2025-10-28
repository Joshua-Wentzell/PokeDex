package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Testing CAPITALS and trailing white space ",
			expected: []string{"testing", "capitals", "and", "trailing", "white", "space"},
		},
		{
			input:    "what is AN Elephant,",
			expected: []string{"what", "is", "an", "elephant,"},
		},
		{
			input:    " what about with only leading space",
			expected: []string{"what", "about", "with", "only", "leading", "space"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual length %d does not match expected length %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("got %q, expected %q", word, expectedWord)
			}
		}
	}
}
