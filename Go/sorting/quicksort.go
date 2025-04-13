package sorting

func median3(s []int) int {
	first := 0
	last := len(s) - 1
	center := (first + last) / 2

	var median int
	if s[first] >= s[last] {
		median = last
	} else {
		median = first
	}

	if s[median] < s[center] {
		median = center
	}

	return median
}

func QuickSort(s []int) {
	if len(s) > 3 {
		pivotIndex := median3(s)
		pivot := s[pivotIndex]
		swap(s, pivotIndex, len(s)-1) // hide pivot

		i, j := 0, len(s)-2
		for {
			for s[i] < pivot {
				i++
			}
			for s[j] > pivot {
				j--
			}
			if i < j {
				swap(s, i, j)
				i++
				j--
			} else {
				break
			}
		}

		swap(s, i, len(s)-1) // restore pivot
		QuickSort(s[:i])
		QuickSort(s[i+1:])
	} else {
		MergeSort(s)
	}
}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}
