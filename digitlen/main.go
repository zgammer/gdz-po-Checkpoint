package main

import (
	"fmt"
	"piscinego1"
)

func main() {
	fmt.Println(piscinego1.DigitLen(100, 10))
	fmt.Println(piscinego1.DigitLen(100, 2))
	fmt.Println(piscinego1.DigitLen(-100, 16))
	fmt.Println(piscinego1.DigitLen(100, -1))
}
