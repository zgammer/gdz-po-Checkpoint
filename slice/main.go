package main

import (
	"fmt"
	"piscinego1"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", piscinego1.Slice(a, 1))
	fmt.Printf("%#v\n", piscinego1.Slice(a, 2, 4))
	fmt.Printf("%#v\n", piscinego1.Slice(a, -3))
	fmt.Printf("%#v\n", piscinego1.Slice(a, -2, -1))
	fmt.Printf("%#v\n", piscinego1.Slice(a, 2, 0))
}
