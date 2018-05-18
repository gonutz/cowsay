package main

import "testing"

func TestSplit(t *testing.T) {
	var tests = []struct {
		have  string
		lines []string
		width int
	}{
		{"", []string{""}, 0},
		{"abc def", []string{"abc def"}, 7},
		{
			"01234567890123456789012345678901234567890123456789012345678901234567890123456789",
			[]string{
				"01234567890123456789012345678901234567890123456789012345678901234567890123",
				"456789",
			},
			74,
		},
		{
			"This text has multiple words in it, seprated by spaces. Spaces are used to split the lines.",
			[]string{
				"This text has multiple words in it, seprated by spaces. Spaces are used",
				"to split the lines.",
			},
			71,
		},
	}

	for _, test := range tests {
		lines, width := splitIntoLines(test.have)
		if !stringsEqual(lines, test.lines) || width != test.width {
			t.Error(width, lines)
		}
	}
}

func stringsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}