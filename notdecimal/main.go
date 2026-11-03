package main

import (
	"fmt"
	"piscinego1"
)

func main() {
	fmt.Print(piscinego1.NotDecimal("0.1"))
	fmt.Print(piscinego1.NotDecimal("174.2"))
	fmt.Print(piscinego1.NotDecimal("0.1255"))
	fmt.Print(piscinego1.NotDecimal("1.20525856"))
	fmt.Print(piscinego1.NotDecimal("-0.0f00d00"))
	fmt.Print(piscinego1.NotDecimal(""))
	fmt.Print(piscinego1.NotDecimal("-19.525856"))
	fmt.Print(piscinego1.NotDecimal("1952"))
}
