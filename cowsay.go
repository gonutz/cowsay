package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func splitIntoLines(s string) (lines []string, maxLineW int) {
	const max = 80 - 6
	runes := []rune(s)
	n := 0
	for _, r := range runes {
		if r == '\n' || unicode.IsGraphic(r) {
			runes[n] = r
			n++
		}
	}
	runes = runes[:n]
	var runeLines [][]rune
	runeLines = append(runeLines, runes)
	wasSplit := true
	for wasSplit {
		n := len(runeLines) - 1
		line := runeLines[n]
		wasSplit = false
		lineBreak := strings.Index(string(line[:min(len(line), max)]), "\n")
		if lineBreak >= 0 {
			runeLines = append(runeLines, line[lineBreak+1:])
			runeLines[n] = line[:lineBreak]
			wasSplit = true
		} else if len(line) > max {
			splitLeft, splitRight := max, max
			for i := max - 1; i >= 1; i-- {
				if line[i] == ' ' {
					splitLeft, splitRight = i+1, i
					break
				}
			}
			runeLines = append(runeLines, line[splitLeft:])
			runeLines[n] = line[:splitRight]
			wasSplit = true
		}
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
