package piscinego1

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	result := 0
	for n != 0 {
		n /= base
		result++
	}
	return result
}
