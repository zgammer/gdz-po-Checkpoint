package main

import (
	"fmt"
	"piscinego1"
)

func main() {
	fmt.Print(piscinego1.FirstWord("hello there"))
	fmt.Print(piscinego1.FirstWord(""))
	fmt.Print(piscinego1.FirstWord("hello   .........  bye"))
}
