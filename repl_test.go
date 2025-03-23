package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "test",
			expected: []string{"test"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		expLen := len(c.expected)
		actLen := len(actual)

		if expLen != actLen {
			t.Errorf("Expected len = %d, actual len = %d", expLen, actLen)
		}

		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("exp = %s, actual = %s", expectedWord, word)
			}

			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
