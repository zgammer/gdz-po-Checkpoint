package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func Atoi(s string) int {
	result := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			break
		}
		result = result*10 + int(char-'0')
	}
	return result
}
func Printl(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
func stringer(n int) string {
	var result string
	if n < 10 {
		return string('0' + n)
	}
	nn := 0
	for i := 10; i <= n; i *= 10 {
		nn = i
	}
	for n != 0 {
		result += string('0' + n/nn)
		n %= nn
		nn /= 10
	}
	return result
}
func main() {
	primenum := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	lb := os.Args[1:]
	var result string
	num := Atoi(lb[0])
	if len(lb) == 1 && num >= 2 {
		if IsPrime(num) {
			Printl(lb[0])
			return
		}
		i := 0
		for true {
			if num < 2 {
				break
			}
			if num%primenum[i] == 0 {
				if len(result) > 0 {
					result += "*"
				}
				result += stringer(primenum[i])
				num /= primenum[i]
			} else {
				i++
			}
		}

	} else {
		return
	}
	Printl(result)
}
