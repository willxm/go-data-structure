package quicksort

func quickSort(source []int, l, u int) []int {
	if l < u {
		m := partition(source, l, u)
		quickSort(source, l, m-1)
		quickSort(source, m, u)
	}
	return source
}

func partition(source []int, l, u int) int {
	var (
		pivot = source[l]
		left  = l
		right = l + 1
	)
	for ; right < u; right++ {
		if source[right] <= pivot {
			left++
			source[left], source[right] = source[right], source[left]
		}
	}
	source[l], source[left] = source[left], source[l]
	return left + 1
}
