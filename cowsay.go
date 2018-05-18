package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var say string
	if len(os.Args) == 1 {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}
		say = string(input)
	} else {
		say = strings.Join(os.Args[1:], " ")
	}

	lines, width := splitIntoLines(say)

	fmt.Print("  ", strings.Repeat("-", width+2), "\n")
	for i := range lines {
		if i == 0 {
			if len(lines) == 1 {
				fmt.Print(" < ")
			} else {
				fmt.Print(" / ")
			}
		} else if i == len(lines)-1 {
			fmt.Print(` \ `)
		} else {
			fmt.Print(" | ")
		}

		fmt.Print(lines[i], strings.Repeat(" ", width+1-len([]rune(lines[i]))))

		if i == 0 {
			if len(lines) == 1 {
				fmt.Print(">")
			} else {
				fmt.Print(`\`)
			}
		} else if i == len(lines)-1 {
			fmt.Print("/")
		} else {
			fmt.Print("|")
		}
		fmt.Print("\n")
	}
	fmt.Print("  ", strings.Repeat("-", width+2), "\n")
	spaces := strings.Repeat(" ", 4+(width-3)/2)
	fmt.Println(spaces + `\   ^--^           `)
	fmt.Println(spaces + ` \  (oo)\_______    `)
	fmt.Println(spaces + `    (__)\       )\/\`)
	fmt.Println(spaces + `        ||----w |   `)
	fmt.Println(spaces + `        ||     ||   `)
}

func splitIntoLines(s string) (lines []string, maxLineW int) {
	const max = 80 - 6
	runes := []rune(s)
	var runeLines [][]rune
	runeLines = append(runeLines, runes)
	for {
		n := len(runeLines) - 1
		if len(runeLines[n]) <= max {
			break
		}
		splitLeft, splitRight := max, max
		for i := max - 1; i >= 1; i-- {
			if runeLines[n][i] == ' ' {
				splitLeft, splitRight = i+1, i
				break
			}
		}
		runeLines = append(runeLines, runeLines[n][splitLeft:])
		runeLines[n] = runeLines[n][:splitRight]
	}
	for _, runeLine := range runeLines {
		if len(runeLine) > maxLineW {
			maxLineW = len(runeLine)
		}
	}
	lines = make([]string, len(runeLines))
	for i := range lines {
		lines[i] = string(runeLines[i])
	}
	return
}
