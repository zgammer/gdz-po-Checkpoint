package main

import (
	"fmt"
	"os"
)

func IsPrime(nb int) bool {
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
func Atoi(s string) int {
	result := 0
	for _, char := range s {
		result = result*10 + int(char-'0')
	}
	return result
}
func main() {
	result := 0
	lb := os.Args[1:]
	num := Atoi(lb[0])
	if len(lb) == 1 && num >= 2 {
		for i := 2; i <= num; i++ {
			if IsPrime(i) {
				result += i
			}
		}
	}
	fmt.Println(result)
}
