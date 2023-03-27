package sorting

func HeapSort(s []int) {
	if len(s) < 2 {
		return
	}
	for i := len(s) / 2; i >= 0; i-- {
		percolateDown(s, i, len(s))
	}

	for i := len(s) - 1; i > 0; i-- {
		swap(s, 0, i)
		percolateDown(s, 0, i)
	}
}

func percolateDown(s []int, i, n int) {
	tmp := s[i]
	for child := leftChild(i); child < n; i, child = child, leftChild(i) {
		if child != n-1 && s[child] < s[child+1] {
			child++
		}
		if tmp > s[child] {
			break
		}
		s[i] = s[child]
	}
	s[i] = tmp
}

func leftChild(i int) int {
	return 2*i + 1
}
