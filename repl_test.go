package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " help\nexit\n",
			expected: []string{"help", "exit"},
		},
		{
			input:    " Balasur Salamander Pikachu",
			expected: []string{"balasur", "salamander", "pikachu"},
		},
	}
	for _, c := range cases {
		out := cleanInput(c.input)
		if len(out) != len(c.expected) {
			t.Errorf("Expected Lenght: %d, OutPut Lenght: %d", len(out), len(c.expected))
		}
		for i := range out {
			actual := out[i]
			expected := c.expected[i]
			if actual != expected {
				t.Errorf("Expected: %s, Got: %s", expected, actual)
			}
		}
	}
}
