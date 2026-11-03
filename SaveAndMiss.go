package piscinego1

func SaveAndMiss(arg string, num int) string {
	result := ""

	if num <= 0 {
		return arg
	}
	i := 0
	j := num
	for j < len(arg) {
		result += string(arg[i:j])
		i += num * 2
		j = i + num
	}
	if j >= len(arg)-1 {
		result += string(arg[i:])
	}
	return result
}
