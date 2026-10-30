package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var w string
	var result string
	lb := os.Args[1:]
	if len(lb) != 1 {
		z01.PrintRune('\n')
		return
	}
	w = lb[0]
	if len(w) > 0 && (rune(w[0]) == '\t' && rune(w[len(w)-1]) == '\t') {
		z01.PrintRune('\n')
		return
	}
	sp := 0
	for in, ch := range w {
		if (in == 0 || in == len(w)-1) && (ch == ' ' || ch == '\t') {
			continue
		}
		if ch != ' ' {
			sp = 0
			result += string(ch)
		}
		if ch == ' ' && sp == 0 {
			sp++
			result += string(ch)
		}
	}
	if rune(result[0]) == ' ' {
		result = result[1:]
	}
	if rune(result[len(result)-1]) == ' ' {
		result = result[:len(result)-1]
	}
	for _, ch1 := range result {
		z01.PrintRune(ch1)
	}
	z01.PrintRune('\n')
}
