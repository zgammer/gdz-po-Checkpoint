package piscinego1

func Slice(a []string, nbrs ...int) []string {
	start := 0
	end := len(a) - 1
	if len(nbrs) >= 1 {
		if nbrs[0] < 0 {
			start = len(a) + nbrs[0]
		} else {
			start = nbrs[0]
		}
	}
	if len(nbrs) == 2 {
		if nbrs[1] < 0 {
			end = len(a) + nbrs[1]
		} else {
			end = nbrs[1]
		}
	}
	if start > end {
		return []string(nil)
	}
	return a[start:end]
}
