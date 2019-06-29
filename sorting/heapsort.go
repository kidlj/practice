package sorting

// Heapsort sorts an int array based on heap tree algorithm.
func Heapsort(a []int) {
	if len(a) < 2 {
		return
	}
	maxIndex := len(a) - 1
	for i := (maxIndex - 1) / 2; i >= 0; i-- {
		siftDown(a, i, maxIndex)
	}
	for {
		a[0], a[maxIndex] = a[maxIndex], a[0]
		maxIndex--
		if maxIndex <= 0 {
			return
		}
		siftDown(a, 0, maxIndex)
	}
}

func siftDown(a []int, i, maxIndex int) {
	tmp := a[i]
	for j := 2*i + 1; j <= maxIndex; j = 2*i + 1 {
		if j < maxIndex && a[j] < a[j+1] {
			j++
		}
		if tmp > a[j] {
			break
		}
		a[i], i = a[j], j
	}
	a[i] = tmp
}
