package piscinego1

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	if a == b {
		return a
	}
	if a > b {
		for i := b; i > 0; i-- {
			if a%i == 0 && b%i == 0 {
				return i
			}
		}
	}
	return Gcd(b, a)
}
