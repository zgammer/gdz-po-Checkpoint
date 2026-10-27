package piscinego1

func Abort(a, b, c, d, e int) int {
	result := []int{a, b, c, d, e}
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}

		}
	}
	return result[2]
}

// для запуска файла нужно написать в терминал go run .\abort\main.go
