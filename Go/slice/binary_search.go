package slice

func binarySearch(s []int, v int) int {
	low := 0
	high := len(s) - 1
	for low <= high {
		mid := (low + high) / 2
		if s[mid] < v {
			low = mid + 1
		} else if s[mid] > v {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
