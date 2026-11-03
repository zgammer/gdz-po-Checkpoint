package piscinego1

import "fmt"

func Chunk(slice []int, size int) {
	var result [][]int
	preresult := []int{}
	if len(slice) == 0 {
		fmt.Println(result)
		return
	}
	if size <= 0 {
		fmt.Println()
		return
	}
	for i := 1; i <= len(slice); i++ {
		if i%size == 0 {
			preresult = append(preresult, slice[i-1])
			result = append(result, preresult)
			preresult = []int{}
		} else {
			preresult = append(preresult, slice[i-1])
		}
	}
	if len(preresult) > 0 {
		result = append(result, preresult)
	}
	fmt.Println(result)
}
