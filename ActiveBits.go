package piscinego1

func ActiveBits(n int) int {
	result := 0
	for n > 0 {
		if n%2 == 1 {
			result++
		}
		n /= 2
	}
	return result
}

//go run .\ActiveBits\main.go
