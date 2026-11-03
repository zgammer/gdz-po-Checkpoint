package piscinego1

func FirstWord(s string) string {
	result := ""
	if len(s) == 0 {
		return "\n"
	}
	for _, ch := range s {
		if ch == ' ' && len(result) > 0 {
			break
		}
		result += string(ch)
	}
	return result + "\n"
}
