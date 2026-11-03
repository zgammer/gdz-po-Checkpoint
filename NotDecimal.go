package piscinego1

func NotDecimal(dec string) string {
	result := ""
	if len(dec) == 0 {
		return "\n"
	}
	for _, ch := range dec {
		if ch >= '0' && ch <= '9' || ch == '-' {
			result += string(ch)
		} else if ch == '.' {
			continue
		} else {
			return dec + "\n"
		}
	}
	if result[0] == '0' {
		result = result[1:]
	}
	return result + "\n"
}
