package piscinego1

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	if n < 0 {
		result = "-" + Itoa(-n/10) + string(rune(-n%10+'0'))
	}
	for n > 0 {
		result = string(rune(n%10+'0')) + result
		n /= 10
	}
	return result
}
