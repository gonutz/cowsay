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
		{
			"Text with line breaks\n\nis split there.\n",
			[]string{
				"Text with line breaks",
				"",
				"is split there.",
				"",
			},
			len("Text with line breaks"),
		},
		{
			"Windows\r\nline breaks...",
			[]string{
				"Windows",
				"line breaks...",
			},
			len("line breaks..."),
		},
		{
			"Print \ttabs\t as spaces",
			[]string{
				"Print  tabs  as spaces",
			},
			len("Print  tabs  as spaces"),
		},
	}

	for _, test := range tests {
		lines, width := splitIntoLines(test.have)
		if !stringsEqual(lines, test.lines) || width != test.width {
			t.Errorf("%d %#v", width, lines)
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
