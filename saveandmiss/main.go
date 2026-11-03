package main

import (
	"fmt"
	"piscinego1"
)

func main() {
	fmt.Println(piscinego1.SaveAndMiss("123456789", 3))
	fmt.Println(piscinego1.SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(piscinego1.SaveAndMiss("", 3))
	fmt.Println(piscinego1.SaveAndMiss("hello you all ! ", 0))
	fmt.Println(piscinego1.SaveAndMiss("what is your name?", 0))
	fmt.Println(piscinego1.SaveAndMiss("go Exercise Save and Miss", -5))
}
