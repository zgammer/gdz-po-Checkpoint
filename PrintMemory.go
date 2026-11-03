package piscinego1

import (
	"unicode"

	"github.com/01-edu/z01"
)

func Hex(a int) {
	result := [2]rune{'0', '0'}
	hexs := "0123456789abcdef"
	if a < 16 {
		result[1] = rune(hexs[a])
	} else {
		result[0] = rune(hexs[a/16])
		result[1] = rune(hexs[a%16])
	}
	for _, r := range result {
		z01.PrintRune(r)
	}
}

func PrintMemory(arr [10]byte) {
	var result string
	for ind, ch := range arr {
		Hex(int(ch))
		if (ind+1)%4 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
		if unicode.IsGraphic(rune(ch)) {
			result += string(ch)
		} else {
			result += "."
		}
	}
	z01.PrintRune('\n')
	for _, rch := range result {
		z01.PrintRune(rune(rch))
	}
	z01.PrintRune('\n')
}
